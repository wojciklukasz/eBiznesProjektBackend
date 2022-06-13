package main

import (
	"ProjektBackend/api/v1/controllers"
	"ProjektBackend/api/v1/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.Connect()

	var err = godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file!")
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	g := e.Group("/api/v1")
	controllers.GetOauthRouting(g)

	e.Logger.Fatal(e.Start(":8000"))
}
