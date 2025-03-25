package main

import (
	"amber/controllers"
	"amber/db"
	"amber/templates/rendering"
	"github.com/labstack/echo/v4"
	"html/template"
)

const ServerAddress = ":8080"

var Router *echo.Echo

func main() {
	Router = echo.New()
	Router.Renderer = &rendering.Template{Templates: template.Must(template.ParseGlob("templates/**/*.html"))}
	Router.Static("public/", "public")
	setupRoutes()
	if err := db.Connect(Router); err != nil {
		Router.Logger.Fatal("Error connecting do Client")
	}
	Router.Logger.Fatal(Router.Start(ServerAddress))
}

func setupRoutes() {
	controllers.SetupHomeRoutes(Router)
	controllers.SetupSnakeRoutes(Router)
}
