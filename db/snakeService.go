package db

import (
	"amber/schemes"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetSnakes() ([]schemes.Snake, error) {
	var snakes []schemes.Snake
	cursor, err := SnakeCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return []schemes.Snake{}, err
	}

	err = cursor.All(context.TODO(), snakes)
	if err != nil {
		return []schemes.Snake{}, err
	}
	return snakes, nil
}

func GetSnake(id bson.ObjectID) (schemes.Snake, error) {
	var snake schemes.Snake
	err := SnakeCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(snake)
	if err != nil {
		return schemes.Snake{}, err
	}
	return snake, nil
}

func SaveSnake(snake schemes.Snake) (schemes.Snake, error) {
	newSnake, err := SnakeCollection.InsertOne(context.TODO(), snake)
	if err != nil {
		return schemes.Snake{}, err
	}
	snake.ID = newSnake.InsertedID.(bson.ObjectID)
	return snake, nil
}
