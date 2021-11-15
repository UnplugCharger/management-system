package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/byron/rest/database"
	"github.com/byron/rest/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var menuCollection *mongo.Collection= database.OpenCollection(database.Client,"menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx , cancel = context.WithTimeout(context.Background(),100*time.Second)
		result, err := menuCollection.Find(context.TODO(),bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error ocurred while trying to list menus"})
		}
		var allMenus []bson.M
		if err = result.All(ctx , &allMenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK,allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	
		return func(c *gin.Context) {
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			menuId := c.Param("menu_id")
			var menu models.Menu
	
			err := foodCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
	
			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching  the menu "})
			}
			c.JSON(http.StatusOK, menu)
		}
	
}
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        var menu  models.Menu


		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(menu)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}


		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Food_id = menu.ID.Hex()
         

		result , insertErr := menuCollection.InsertOne(ctx,menu)

		if insertErr != nil {
			msg := fmt.Sprintf("Menu Item was not Created")
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
		}
		defer cancel()
		c.JSON(http.StatusOK,result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
