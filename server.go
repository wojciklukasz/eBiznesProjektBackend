package main

import (
	"ProjektBackend/api/v1/controllers"
	"ProjektBackend/api/v1/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	database.Connect()

	e := echo.New()
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"http://localhost:3000", "https://ebiznesprojekt.azurewebsites.net:3000"},
	//}))
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})

	g := e.Group("/api/v1")
	controllers.GetOauthRouting(g)
	controllers.GetCategoriesRouting(g)
	controllers.GetManufacturersRouting(g)
	controllers.GetProductsRouting(g)
	controllers.GetUsersRouting(g)
	controllers.GetOrdersRouting(g)
	controllers.GetPaymentRouting(g)

	e.Logger.Fatal(e.Start(":3051"))
}
