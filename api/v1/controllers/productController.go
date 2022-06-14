package controllers

import (
	"ProjektBackend/api/v1/database"
	"ProjektBackend/api/v1/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var itemNotFoundMessage = "Item not found"

func GetProductsRouting(e *echo.Group) {
	g := e.Group("/product")
	g.GET("", GetProducts)
	g.GET("/:id", GetProduct)
	g.POST("", SaveProduct)
	g.PUT("/:id", UpdateProduct)
	g.DELETE("/:id", DeleteProduct)
}

func GetProducts(c echo.Context) error {
	fmt.Println("\n\nADRES:    " + c.Request().Host + "\n\n")
	var products []models.Product

	result := database.Database.Find(&products)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Items not found")
	}

	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product

	result := database.Database.Find(&product, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	return c.JSON(http.StatusOK, product)
}

func SaveProduct(c echo.Context) error {
	product := new(models.Product)

	err := c.Bind(product)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid body "+err.Error())
	}

	result := database.Database.Create(product)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "Database error "+result.Error.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	result := database.Database.Find(&product, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	values := new(models.Product)
	err := c.Bind(values)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid body "+err.Error())
	}

	product.Name = values.Name
	product.Price = values.Price
	product.CategoryID = values.CategoryID
	product.ManufacturerID = values.ManufacturerID
	product.Description = values.Description
	database.Database.Save(&product)

	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product

	result := database.Database.Delete(&product, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "item deleted"})
}
