package controllers

import (
	"amber/db"
	"amber/schemes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
)

func SetupTerrariumRoutes(router *echo.Echo) {
	router.GET("/terrarium/new", func(c echo.Context) error {
		err := c.Render(http.StatusOK, "add-terrarium-page", "")
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		return nil
	})

	router.POST("/terrarium", func(c echo.Context) error {
		newTerrarium := schemes.TerrariumJson{}
		body := c.Request().Body
		err := json.NewDecoder(body).Decode(&newTerrarium)
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		convertedTerrarium, err := schemes.ConvertToTerrarium(newTerrarium)
		convertedTerrarium.ID = bson.NewObjectID()
		_, err = db.SaveTerrarium(convertedTerrarium)
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		return nil
	})
}
