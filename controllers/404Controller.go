package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Setup404Routes(router *echo.Echo) {
	router.GET("*", func(c echo.Context) error {
		err := c.Render(http.StatusNotFound, "404-page", "")
		if err != nil {
			router.Logger.Error(err)
			return err
		}
		return nil
	})
}
