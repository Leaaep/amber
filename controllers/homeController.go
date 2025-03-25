package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupHomeRoutes(router *echo.Echo) {
	router.GET("/", func(c echo.Context) error {
		//_, err := db.GetSnakes()
		//if err != nil {
		//	return err
		//}

		err := c.Render(http.StatusOK, "home-page", "")
		if err != nil {
			return err
		}
		return nil
	})
}
