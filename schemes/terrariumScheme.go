package schemes

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"strconv"
	"time"
)

type Terrarium struct {
	ID                  bson.ObjectID `bson:"_id"`
	Name                string        `bson:"name"`
	Snakes              []Snake       `bson:"snakes"`
	Length              int64         `bson:"length"`
	Width               int64         `bson:"width"`
	Height              int64         `bson:"height"`
	LastMaintenanceDate time.Time     `bson:"lastMaintenanceDate"`
	MaintenanceInterval int64         `bson:"maintenanceInterval"`
}

type TerrariumJson struct {
	ID                  bson.ObjectID `bson:"_id"`
	Name                string        `bson:"name"`
	Snakes              []SnakeJson   `bson:"snakes"`
	Length              string        `bson:"length"`
	Width               string        `bson:"width"`
	Height              string        `bson:"height"`
	LastMaintenanceDate string        `bson:"lastMaintenanceDate"`
	MaintenanceInterval string        `bson:"maintenanceInterval"`
}

func ConvertToTerrarium(json TerrariumJson) (Terrarium, error) {

	length, err := strconv.ParseInt(json.Length, 0, 0)
	width, err := strconv.ParseInt(json.Width, 0, 0)
	height, err := strconv.ParseInt(json.Height, 0, 0)
	lastMaintenanceDate, err := time.Parse("2006-01-02", json.LastMaintenanceDate)
	maintenanceInterval, err := strconv.ParseInt(json.MaintenanceInterval, 0, 0)

	if err != nil {
		return Terrarium{}, err
	}

	var snakes []Snake
	for _, snake := range json.Snakes {
		convertedSnake, err := ConvertToSnake(snake)
		if err != nil {
			return Terrarium{}, err
		}
		snakes = append(snakes, convertedSnake)
	}

	return Terrarium{
			ID:                  json.ID,
			Name:                json.Name,
			Snakes:              snakes,
			Length:              length,
			Width:               width,
			Height:              height,
			LastMaintenanceDate: lastMaintenanceDate,
			MaintenanceInterval: maintenanceInterval,
		},
		nil
}
