package main

import (
	"ProjektBackend/api/v1/controllers"
	"ProjektBackend/api/v1/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "heh")
	})

	e.GET("/google", func(c echo.Context) error {
		url := controllers.GetLoginURL("google")
		return c.JSON(http.StatusOK, map[string]string{"url": url})
	})

	e.GET("/github", func(c echo.Context) error {
		url := controllers.GetLoginURL("github")
		return c.JSON(http.StatusOK, map[string]string{"url": url})
	})

	e.GET("/auth/google/callback", controllers.HandleGoogleCallback)
	e.GET("/auth/github/callback", controllers.HandleGithubCallback)

	e.Logger.Fatal(e.Start(":8000"))
}
