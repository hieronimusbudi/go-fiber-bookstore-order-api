package mongodb

import (
	"context"
	"log"

	envvar "github.com/hieronimusbudi/go-fiber-bookstore-order-api/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	OrdersCollection *mongo.Collection
	ItemsCollection  *mongo.Collection
	mongoURI         = envvar.MongoURI
	mongoURILocal    = "mongodb://localhost:27017/orders"
)

func InitMongo() {
	clientOptions := options.Client().ApplyURI(mongoURILocal)

	client, clientErr := mongo.Connect(context.TODO(), clientOptions)
	if clientErr != nil {
		panic(clientErr)
	}

	pingErr := client.Ping(context.TODO(), nil)
	if pingErr != nil {
		panic(pingErr)
	}

	log.Println("Connected to MongoDB!")

	OrdersCollection = client.Database("orders").Collection("orders")
	ItemsCollection = client.Database("orders").Collection("items")
}
