package controllers

import (
	"context"
	"fmt"
	"github.com/byron/rest/database"
	"github.com/byron/rest/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing items"})
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders[0])

	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("order_id")
		var order models.Order
		err := foodCollection.FindOne(ctx, bson.M{"order_id": foodId}).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the orders"})
		}
		c.JSON(http.StatusOK, order)
	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table
		var order models.Order

		var updateObj primitive.D

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if order.TableId != nil {
			err := menuCollection.FindOne(ctx, bson.M{"table_id": order.TableId}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("Message:Table was not found ")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return

			}
			updateObj = append(updateObj, bson.E{"menu", order.TableId})
		}
		order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.UpdatedAt})
		upsert := true

		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := orderCollection.UpdateOne(ctx,
			filter,
			bson.D{
				{"$st", updateObj},
			},
			&opt,
		)
		if err != nil {
			msg := fmt.Sprintf("order update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table
		var order models.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(order)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		if order.TableId != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.TableId}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("Table was not found ")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
		}
		order.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.OrderId = order.ID.Hex()

		result, inserErr := orderCollection.InsertOne(ctx, order)

		if inserErr != nil {
			msg := fmt.Sprintf("order  item was not creates")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}

}
func OrderItemCreator(order models.Order) string {
	order.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.OrderId = order.ID.Hex()
	orderCollection.InsertOne(ctx, order)
	defer cancel()

	return order.OrderId
}
