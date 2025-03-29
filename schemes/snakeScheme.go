package schemes

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"strconv"
	"time"
)

type Snake struct {
	ID                   bson.ObjectID `bson:"_id"`
	Name                 string        `bson:"name"`
	Birthdate            time.Time     `bson:"birthdate"`
	LastFeedingDate      time.Time     `bson:"lastFeedingDate"`
	FeedingInterval      int64         `bson:"feedingInterval"`
	WinterBreakStartDate time.Time     `bson:"winterBreakStartDate"`
	WinterBreakDuration  int64         `bson:"winterBreakDuration"`
}

type SnakeJson struct {
	ID                   bson.ObjectID `bson:"_id"`
	Name                 string        `bson:"name"`
	Birthdate            string        `bson:"birthdate"`
	LastFeedingDate      string        `bson:"lastFeedingDate"`
	FeedingInterval      string        `bson:"feedingInterval"`
	WinterBreakStartDate string        `bson:"winterBreakStartDate"`
	WinterBreakDuration  string        `bson:"winterBreakDuration"`
}

func ConvertToSnake(json SnakeJson) (Snake, error) {

	birthdate, err := time.Parse("2006-01-02", json.Birthdate)
	lastFeedingDate, err := time.Parse("2006-01-02", json.LastFeedingDate)
	feedingInterval, err := strconv.ParseInt(json.FeedingInterval, 0, 0)
	winterBreakStartDate, err := time.Parse("2006-01-02", json.WinterBreakStartDate)
	winterBreakDuration, err := strconv.ParseInt(json.WinterBreakDuration, 0, 0)

	if err != nil {
		return Snake{}, err
	}

	return Snake{
		ID:                   json.ID,
		Name:                 json.Name,
		Birthdate:            birthdate,
		LastFeedingDate:      lastFeedingDate,
		FeedingInterval:      feedingInterval,
		WinterBreakStartDate: winterBreakStartDate,
		WinterBreakDuration:  winterBreakDuration,
	}, nil
}
