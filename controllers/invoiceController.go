package controllers

import (
	"context"
	"github.com/byron/rest/database"
	"github.com/byron/rest/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

//Invoice format
type InvoiceViewFormat struct {
	InvoiceId      string
	PaymentMethod  string
	OrderId        string
	PaymentStatus  *string
	PaymentDue     interface{}
	TableNUmber    interface{}
	PaymentDueDate time.Time
	OrderDetails   interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving invoice"})
		}
		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil {
			c.JSON(http.StatusOK, allInvoices)
		}
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")

		var invoice models.Invoice

		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing invoice item"})
		}
		var invoiceView InvoiceViewFormat

		allOrderItems, err := ItemsByOrder(invoice.Order_id)

		invoiceView.OrderId = invoice.Order_id
	}
}
func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
