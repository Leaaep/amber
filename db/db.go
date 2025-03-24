package db

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const MongoUrl = "mongodb://amber-db:27017/"
const DatabaseName = "amber_db"

var Client *mongo.Client
var Database *mongo.Database
var SnakeCollection *mongo.Collection
var TerrariumCollection *mongo.Collection

func Connect(router *echo.Echo) error {
	clientOption := options.Client().ApplyURI(MongoUrl)
	client, err := mongo.Connect(clientOption)
	if err != nil {
		router.Logger.Error(err)
	}

	err = client.Ping(context.Background(), nil)
	router.Logger.Print("Connected to MongoDB")
	Client = client
	return err
}

func setupDB() {
	Database = Client.Database(DatabaseName)
	SnakeCollection = Database.Collection("snake")
	TerrariumCollection = Database.Collection("terrarium")
}
