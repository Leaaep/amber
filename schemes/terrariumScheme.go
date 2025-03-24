package schemes

import "go.mongodb.org/mongo-driver/v2/bson"

type Terrarium struct {
	ID               bson.ObjectID `bson:"_id"`
	Name             string        `bson:"name"`
	Snakes           []Snake       `bson:"snakes"`
	Length           int16         `bson:"length"`
	Width            int16         `bson:"width"`
	Height           int16         `bson:"height"`
	MaintenanceCycle int8          `bson:"maintenanceCycle"`
}

func (t Terrarium) New(name string, snakes []Snake, length int16, width int16, height int16, maintenanceCycle int8) Terrarium {
	return Terrarium{
		ID:               bson.ObjectID{},
		Name:             name,
		Snakes:           snakes,
		Length:           length,
		Width:            width,
		Height:           height,
		MaintenanceCycle: maintenanceCycle,
	}
}
