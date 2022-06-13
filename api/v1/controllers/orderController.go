package controllers

import (
	"ProjektBackend/api/v1/database"
	"ProjektBackend/api/v1/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetOrdersRouting(e *echo.Group) {
	g := e.Group("/order")
	g.GET("", GetOrders)
	g.GET("/:id", GetOrder)
	g.POST("", SaveOrder)
	g.PUT("/:id", UpdateOrder)
	g.DELETE("/:id", DeleteOrder)
}

func GetOrders(c echo.Context) error {
	var orders []models.Order

	result := database.Database.Find(&orders)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Items not found")
	}

	return c.JSON(http.StatusOK, orders)
}

func GetOrder(c echo.Context) error {
	id := c.Param("id")
	var order models.Order

	result := database.Database.Find(&order, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Item not found")
	}

	return c.JSON(http.StatusOK, order)
}

func SaveOrder(c echo.Context) error {
	order := new(models.Order)

	err := c.Bind(order)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid body "+err.Error())
	}

	result := database.Database.Create(order)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "Database error "+result.Error.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func UpdateOrder(c echo.Context) error {
	id := c.Param("id")
	var order models.Order
	result := database.Database.Find(&order, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Item not found")
	}

	values := new(models.Order)
	err := c.Bind(values)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid body "+err.Error())
	}

	order.Date = values.Date
	order.CustomerID = values.CustomerID
	order.Total = values.Total
	database.Database.Save(&order)

	return c.JSON(http.StatusOK, order)
}

func DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	var order models.Order

	result := database.Database.Delete(&order, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Item not found")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "item deleted"})
}
