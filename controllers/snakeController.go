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
	router.GET("/snake/:id", func(c echo.Context) error {
		id, err := bson.ObjectIDFromHex(c.Param("id"))
		snake, err := db.GetSnake(id)
		if err != nil {
			return err
		}

		err = c.Render(http.StatusOK, "snake-card-component", snake)
		if err != nil {
			return err
		}
		return nil
	})

	router.GET("/snake/new", func(c echo.Context) error {
		err := c.Render(http.StatusOK, "add-snake-page", "")
		if err != nil {
			return err
		}
		return nil
	})
	router.POST("/snake", func(c echo.Context) error {
		newSnake := schemes.Snake{}
		err := json.NewDecoder(c.Request().Body).Decode(&newSnake)
		if err != nil {
			return err
		}
		_, err = db.SaveSnake(newSnake)
		if err != nil {
			return err
		}
		return nil
	})
}
