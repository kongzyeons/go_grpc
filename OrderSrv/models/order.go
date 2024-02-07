package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   uint               `bson:"user_id" json:"user_id"`
	Products []struct {
		ProductID uint `bson:"product_id" json:"product_id"`
		Amount    int  `bson:"amount" json:"amount"`
	} `bson:"products" json:"products"`
	StatusOrder string    `bson:"status_order" json:"status_order"`
	CreateTime  time.Time `bson:"createTime" json:"create_time"`
	UpdateTime  time.Time `bson:"updateTime" json:"update_time"`
}
