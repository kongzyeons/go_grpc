package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   uint               `bson:"user_id"`
	Products []struct {
		ProducrtID uint `bson:"product_id"`
		Amount     int  `bson:"amount"`
	} `bson:"producrs"`
	StatusOrder string    `bson:"status_order"`
	CreateTime  time.Time `bson:"createTime"`
	UpdateTime  time.Time `bson:"updateTime"`
}
