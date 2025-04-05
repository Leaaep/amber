package main

import (
	"amber/controllers"
	"amber/db"
	"amber/rendering"
	"github.com/labstack/echo/v4"
	"html/template"
	"time"
)

const ServerAddress = ":8080"

var router *echo.Echo

func main() {
	router = echo.New()

	templates := template.New("").Funcs(getFuncMap())
	templates = template.Must(templates.ParseGlob("templates/**/*.html"))

	router.Renderer = &rendering.Template{Templates: templates}
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

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"addDays": func(days int64, t time.Time) time.Time {
			return t.AddDate(0, 0, int(days))
		},
		"timeNow": func(t time.Time) string {
			t = time.Now()
			return t.Format("2006-01-02")
		},
	}
}
