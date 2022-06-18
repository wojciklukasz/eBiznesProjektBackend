package controllers

import (
	"ProjektBackend/api/v1/database"
	"ProjektBackend/api/v1/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	return c.JSON(http.StatusOK, order)
}

func SaveOrder(c echo.Context) error {
	order := new(models.Order)

	data := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid body"})
	}

	itemOrders := data["items"].(map[string]interface{})
	var items []models.ItemOrder
	for product, count := range itemOrders {
		id, _ := strconv.ParseInt(product, 10, 32)
		count, _ := strconv.ParseInt(count.(string), 10, 32)
		result := SaveItemOrder(int(id), int(count))
		items = append(items, result)
	}

	order.Name = data["name"].(string)
	order.Surname = data["surname"].(string)
	order.Email = data["email"].(string)
	order.Road = data["road"].(string)
	order.Nr = data["nr"].(string)
	order.Code = data["code"].(string)
	order.City = data["city"].(string)
	order.Phone = data["phone"].(string)
	order.Items = items

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
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	values := new(models.Order)
	err := c.Bind(values)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid body "+err.Error())
	}

	order.Total = values.Total
	database.Database.Save(&order)

	return c.JSON(http.StatusOK, order)
}

func DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	var order models.Order

	result := database.Database.Delete(&order, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "item deleted"})
}
