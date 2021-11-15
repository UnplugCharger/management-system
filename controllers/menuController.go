package controllers

import (
	"context"

	"github.com/byron/rest/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var menuCollection *mongo.Collection= database.OpenCollection(database.Client,"menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := menuCollection.Find(context.TODO(),bson.M{})
		defer cancel()
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
