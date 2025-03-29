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

	router.GET("/terrarium/:id", func(c echo.Context) error {
		hexId := c.Param("id")
		router.Logger.Print(hexId)
		router.Logger.Print(bson.ObjectIDFromHex(hexId))
		terrarium, err := db.GetTerrarium(hexId)
		if err != nil {
			router.Logger.Error(err)
			err := c.Render(http.StatusNotFound, "404-page", "")
			if err != nil {
				return err
			}
			return nil
		}

		err = c.Render(http.StatusOK, "terrarium-page", terrarium)
		if err != nil {
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

	/* --- SNAKE ROUTES --- */
	router.GET("/terrarium/:terrariumID/snake/new", func(c echo.Context) error {
		hexId := c.Param("terrariumID")
		terrarium, err := db.GetTerrarium(hexId)
		if err != nil {
			router.Logger.Error(err)
			err := c.Render(http.StatusNotFound, "404-page", "")
			if err != nil {
				return err
			}
			return nil
		}

		err = c.Render(http.StatusOK, "add-snake-page", terrarium)
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		return nil
	})

	router.GET("/terrarium/:terrariumID/snake/:snakeID", func(c echo.Context) error {
		terrariumID := c.Param("terrariumID")
		snakeID := c.Param("snakeID")

		terrarium, err := db.GetTerrarium(terrariumID)
		if err != nil {
			router.Logger.Error(err)
			err := c.Render(http.StatusNotFound, "404-page", "")
			if err != nil {
				return err
			}
			return nil
		}

		for _, snake := range terrarium.Snakes {
			if snake.ID.Hex() == snakeID {
				err = c.Render(http.StatusOK, "snake-page", snake)
				if err != nil {
					return err
				}
				return nil
			}
		}
		return nil
	})
	router.POST("/terrarium/:terrariumID/snake", func(c echo.Context) error {
		terrariumID := c.Param("terrariumID")

		newSnake := schemes.SnakeJson{}
		body := c.Request().Body
		err := json.NewDecoder(body).Decode(&newSnake)
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		convertedSnake, err := schemes.ConvertToSnake(newSnake)
		if err != nil {
			router.Logger.Error(err)
			return err
		}

		convertedSnake.ID = bson.NewObjectID()

		terrarium, err := db.GetTerrarium(terrariumID)
		if err != nil {
			router.Logger.Error(err)
			return err
		}

		terrarium.Snakes = append(terrarium.Snakes, convertedSnake)
		err = db.UpdateTerrarium(terrarium, terrarium.ID)
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		return nil
	})
}
