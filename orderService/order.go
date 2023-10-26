package order

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order struct {
	ID           string             `json:"_id" bson:"_id"`
	ProductId    primitive.ObjectID `json:"productId" bson:"productId"`
	Quantity     int                `json:"quantity" bson:"quantity"`
	IsPremium    bool               `json:"isPremium" bson:"isPremium"`
	OrderValue   float64            `json:"orderValue" bson:"orderValue"`
	DispatchDate time.Time          `json:"dispatchDate,omitempty" bson:"dispatchDate,omitempty"`
	Status       string             `json:"status" bson:"status"`
}

var orderCollection *mongo.Collection

func InitOrderCollection(client *mongo.Client, dbName, collectionName string) {
	orderCollection = client.Database(dbName).Collection(collectionName)
}

func CreateOrder(order *Order) error {
	_, err := orderCollection.InsertOne(context.TODO(), order)
	return err
}

func UpdateOrderStatus(orderID, status string) error {
	_, err := orderCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": orderID},
		bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: status}}}},
	)
	return err
}
