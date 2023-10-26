package product

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id" `
	Name         string             `json:"name" bson:"name"`
	Availability int                `json:"availability" bson:"availability"`
	Price        float64            `json:"price" bson:"price"`
	Category     string             `json:"category" bson:"category"`
}

var productCollection *mongo.Collection

func InitProductCollection(client *mongo.Client, dbName, collectionName string) {
	productCollection = client.Database(dbName).Collection(collectionName)
}

func GetProducts() ([]Product, error) {
	var products []Product
	cursor, err := productCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(productId primitive.ObjectID) (*Product, error) {
	var product Product
	err := productCollection.FindOne(context.TODO(), bson.M{"_id": productId}, nil).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &product, nil
}

func UpdateProduct(product *Product) error {
	_, err := productCollection.UpdateOne(context.TODO(), bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
