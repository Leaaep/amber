package main

import (
	"amber/db"
	"github.com/labstack/echo/v4"
)

const ServerAddress = ":8080"

var Router *echo.Echo

func main() {
	Router = echo.New()
	if err := db.Connect(Router); err != nil {
		Router.Logger.Fatal("Error connecting do Client")
	}
	Router.Logger.Fatal(Router.Start(ServerAddress))
}
