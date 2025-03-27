package db

import (
	"amber/schemes"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetTerrariums() ([]schemes.Terrarium, error) {
	var terrariums []schemes.Terrarium
	cursor, err := TerrariumCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		return []schemes.Terrarium{}, err
	}

	err = cursor.All(context.TODO(), terrariums)
	if err != nil {
		return []schemes.Terrarium{}, err
	}
	return terrariums, nil
}

func GetTerrarium(id bson.ObjectID) (schemes.Terrarium, error) {
	var terrarium schemes.Terrarium
	err := TerrariumCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(terrarium)
	if err != nil {
		return schemes.Terrarium{}, err
	}
	return terrarium, nil

}

func SaveTerrarium(terrarium schemes.Terrarium) (schemes.Terrarium, error) {
	newTerrarium, err := TerrariumCollection.InsertOne(context.TODO(), terrarium)
	if err != nil {
		return schemes.Terrarium{}, err
	}
	terrarium.ID = newTerrarium.InsertedID.(bson.ObjectID)
	return terrarium, nil

}
