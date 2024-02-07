package models

import (
	"time"
)

type Order struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	UserID   uint   `bson:"user_id" json:"user_id"`
	Products []struct {
		ProductID uint `bson:"product_id" json:"product_id"`
		Amount    int  `bson:"amount" json:"amount"`
	} `bson:"products" json:"products"`
	StatusOrder string    `bson:"status_order" json:"status_order"`
	CreateTime  time.Time `bson:"createTime" json:"create_time"`
	UpdateTime  time.Time `bson:"updateTime" json:"update_time"`
}

type CreateOrderReq struct {
	ProductID int `json:"product_id" binding:"required"`
	Amount    int `json:"amount" binding:"required"`
}

type OrderProduct struct {
	Product Product `json:"product"`
	Amount  int     `json:"amount"`
}

type GetOrderByUserRes struct {
	OrderID     string         `json:"order_id"`
	UserID      uint32         `json:"user_id"`
	StatusOrder string         `json:"status_order"`
	Products    []OrderProduct `json:"products"`
	CreateTime  time.Time      `json:"create_time"`
	UpdateTime  time.Time      `json:"update_time"`
}

type AddProductReq struct {
	ProductID int `json:"product_id" binding:"required"`
	Amount    int `json:"amount" binding:"required"`
}
