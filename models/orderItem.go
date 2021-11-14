package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)







type OrderItem struct {
	ID  primitive.ObjectID  `bson:"_id"`
	Quantity   *string  `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Unit_price *float64  `json:"unit_price" validate:"required"`
	Created_at time.Time   
	Updated_at time.Time 
	Food_id  *string 
	Order_item_id  string
	Order_id  string
}