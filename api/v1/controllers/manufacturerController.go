package controllers

import (
	"ProjektBackend/api/v1/database"
	"ProjektBackend/api/v1/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetManufacturersRouting(e *echo.Group) {
	g := e.Group("/manufacturer")
	g.GET("", GetManufacturers)
	g.GET("/:id", GetManufacturer)
	g.POST("", SaveManufacturer)
	g.PUT("/:id", UpdateManufacturer)
	g.DELETE("/:id", DeleteManufacturer)
}

func GetManufacturers(c echo.Context) error {
	var manufacturers []models.Manufacturer

	result := database.Database.Find(&manufacturers)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Items not found")
	}

	return c.JSON(http.StatusOK, manufacturers)
}

func GetManufacturer(c echo.Context) error {
	id := c.Param("id")
	var manufacturer models.Manufacturer

	result := database.Database.Find(&manufacturer, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	return c.JSON(http.StatusOK, manufacturer)
}

func SaveManufacturer(c echo.Context) error {
	manufacturer := new(models.Manufacturer)

	err := c.Bind(manufacturer)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid body "+err.Error())
	}

	result := database.Database.Create(manufacturer)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "Database error "+result.Error.Error())
	}

	return c.JSON(http.StatusOK, manufacturer)
}

func UpdateManufacturer(c echo.Context) error {
	id := c.Param("id")
	var manufacturer models.Manufacturer
	result := database.Database.Find(&manufacturer, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	values := new(models.Manufacturer)
	err := c.Bind(values)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid body "+err.Error())
	}

	manufacturer.Name = values.Name
	manufacturer.Description = values.Description
	database.Database.Save(&manufacturer)

	return c.JSON(http.StatusOK, manufacturer)
}

func DeleteManufacturer(c echo.Context) error {
	id := c.Param("id")
	var manufacturer models.Manufacturer

	result := database.Database.Delete(&manufacturer, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "item deleted"})
}
