package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(url string) *mongo.Client {
	connectionDB, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = connectionDB.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect mongodb :", url)
	return connectionDB

}
