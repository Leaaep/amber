package controllers

import (
	"amber/db"
	"amber/schemes"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
)

func GetSnake(id bson.ObjectID) (schemes.Snake, error) {
	var snake schemes.Snake
	err := db.SnakeCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(snake)
	if err != nil {
		log.Fatal(err)
		return schemes.Snake{}, err
	}
	return snake, nil
}

func GetSnakes() ([]schemes.Snake, error) {
	var snakes []schemes.Snake
	cursor, err := db.SnakeCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return []schemes.Snake{}, err
	}

	err = cursor.All(context.TODO(), snakes)
	if err != nil {
		log.Fatal(err)
		return []schemes.Snake{}, err
	}
	return snakes, nil
}

func AddSnake(snake schemes.Snake) (schemes.Snake, error) {
	newSnake, err := db.SnakeCollection.InsertOne(context.TODO(), snake)
	if err != nil {
		return schemes.Snake{}, err
	}
	snake.ID = bson.ObjectID(newSnake.InsertedID.(primitive.ObjectID))
	if snake.ID.IsZero() {
		return schemes.Snake{}, errors.New("no ID to new snakes assigned")
	}
	return snake, nil
}
