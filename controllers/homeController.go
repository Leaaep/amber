package controllers

import (
	"amber/db"
	"amber/schemes"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
	"time"
)

type HomePage struct {
	ID                  bson.ObjectID   `bson:"_id"`
	Name                string          `bson:"name"`
	Snakes              []schemes.Snake `bson:"snakes"`
	Length              int64           `bson:"length"`
	Width               int64           `bson:"width"`
	Height              int64           `bson:"height"`
	LastMaintenanceDate time.Time       `bson:"lastMaintenanceDate"`
	MaintenanceInterval int64           `bson:"maintenanceInterval"`
}

func SetupHomeRoutes(router *echo.Echo) {

	router.GET("/", func(c echo.Context) error {
		terrariums, err := db.GetTerrariums()
		if err != nil {
			router.Logger.Error(err)
			return err
		}

		err = c.Render(http.StatusOK, "home-page", terrariums)
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		return nil
	})
}
