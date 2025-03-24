package schemes

import "go.mongodb.org/mongo-driver/v2/bson"

type Snake struct {
	ID               bson.ObjectID `bson:"_id"`
	Name             string        `bson:"name"`
	Birthdate        bson.DateTime `bson:"birthdate"`
	FeedCycle        int8          `bson:"feedCycle"`
	WinterBreakCycle int8          `bson:"winterBreakCycle"`
}

func (s Snake) New(name string, birthdate bson.DateTime, feedCycle int8, winterBreakCycle int8) Snake {
	return Snake{
		ID:               bson.ObjectID{},
		Name:             name,
		Birthdate:        birthdate,
		FeedCycle:        feedCycle,
		WinterBreakCycle: winterBreakCycle,
	}
}
