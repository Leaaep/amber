package db

import (
	"amber/schemes"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetTerrariums() ([]schemes.Terrarium, error) {
	var terrariums []schemes.Terrarium
	cursor, err := TerrariumCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return terrariums, err
	}

	err = cursor.All(context.Background(), &terrariums)
	if err != nil {
		return terrariums, err
	}
	return terrariums, nil
}

func GetTerrarium(id string) (schemes.Terrarium, error) {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return schemes.Terrarium{}, err
	}

	var terrarium schemes.Terrarium
	err = TerrariumCollection.FindOne(context.Background(), bson.D{{"_id", objectId}}).Decode(&terrarium)
	if err != nil {
		return schemes.Terrarium{}, err
	}
	return terrarium, nil

}

func SaveTerrarium(terrarium schemes.Terrarium) (schemes.Terrarium, error) {
	_, err := TerrariumCollection.InsertOne(context.Background(), terrarium)
	if err != nil {
		return schemes.Terrarium{}, err
	}
	return terrarium, nil

}

func UpdateTerrarium(terrarium schemes.Terrarium, id bson.ObjectID) error {
	_, err := TerrariumCollection.UpdateByID(context.Background(), id, bson.D{{"$set", bson.D{{"snakes", terrarium.Snakes}}}})
	if err != nil {
		return err
	}
	return nil
}
