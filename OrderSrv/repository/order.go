package repository

import (
	"context"
	"log"
	"my-package/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepo interface {
	Create(order models.Order) error
	GetQuery(order models.Order) (orders []models.Order, err error)
	Update(order models.Order) error
	Delete(id string) error
}

type orderRepo struct {
	db *mongo.Client
}

func NewOrderRepo(db *mongo.Client) OrderRepo {
	return orderRepo{db}
}

func (obj orderRepo) Create(order models.Order) error {
	collection := obj.db.Database("Order").Collection("orders")
	_, err := collection.InsertOne(context.TODO(), order)
	if err != nil {
		log.Println("Error inserting order:", err)
		return err
	}
	return nil
}

// GetQuery retrieves orders based on the provided query parameters.
func (obj orderRepo) GetQuery(query models.Order) (orders []models.Order, err error) {
	collection := obj.db.Database("Order").Collection("orders")

	var pipeline []bson.M
	var defaultOrder models.Order
	if query.ID != defaultOrder.ID {
		pipeline = append(pipeline, bson.M{
			"$match": bson.M{"_id": query.ID},
		})
	}
	if query.UserID != 0 {
		pipeline = append(pipeline, bson.M{
			"$match": bson.M{"user_id": query.UserID},
		})
	}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// Update modifies an existing order in the database.
func (obj orderRepo) Update(order models.Order) error {
	collection := obj.db.Database("Order").Collection("orders")

	filter := bson.M{"_id": order.ID} // Assuming Order struct has an ID field

	update := bson.M{"$set": order}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating order:", err)
		return err
	}
	return nil
}

// Delete removes an order from the database based on its ID.
func (obj orderRepo) Delete(id string) error {
	collection := obj.db.Database("Order").Collection("orders")

	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error deleting order:", err)
		return err
	}
	return nil
}
