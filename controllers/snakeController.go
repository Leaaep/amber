package controllers

import (
	"amber/db"
	"amber/schemes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
)

func SetupSnakeRoutes(router *echo.Echo) {
	router.GET("/snake/new", func(c echo.Context) error {
		err := c.Render(http.StatusOK, "add-snake-page", "")
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		return nil
	})
	router.POST("/snake", func(c echo.Context) error {
		newSnake := schemes.SnakeJson{}
		body := c.Request().Body
		err := json.NewDecoder(body).Decode(&newSnake)
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		convertedSnake, err := schemes.ConvertToSnake(newSnake)
		convertedSnake.ID = bson.NewObjectID()
		_, err = db.SaveSnake(convertedSnake)
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		return nil
	})
}
