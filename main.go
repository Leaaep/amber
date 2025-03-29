package main

import (
	"amber/controllers"
	"amber/db"
	"amber/rendering"
	"github.com/labstack/echo/v4"
	"html/template"
)

const ServerAddress = ":8080"

var router *echo.Echo

func main() {
	router = echo.New()

	router.Renderer = &rendering.Template{Templates: template.Must(template.ParseGlob("templates/**/*.html"))}
	router.Static("public/", "public")

	setupRoutes()
	if err := db.Connect(router); err != nil {
		router.Logger.Fatal("Error connecting do Client")
	}
	router.Logger.Fatal(router.Start(ServerAddress))
}

func setupRoutes() {
	controllers.SetupHomeRoutes(router)
	controllers.SetupTerrariumRoutes(router)
	controllers.Setup404Routes(router)
}
