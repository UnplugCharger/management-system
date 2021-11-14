package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



type User struct {
	ID primitive.ObjectID  `bson:"_id"`
	First_name   *string 
	Last_name  *string 
	Password *string 
	Email *string 
	Avatar *string 
	Phone *string 
	Token *string 
	Refresh_token *string 
	Created_at time.Time
	Updated_at time.Time
	User_id  string
} 