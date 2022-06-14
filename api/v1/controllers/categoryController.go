package controllers

import (
	"ProjektBackend/api/v1/database"
	"ProjektBackend/api/v1/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCategoriesRouting(e *echo.Group) {
	g := e.Group("/category")
	g.GET("", GetCategories)
	g.GET("/:id", GetCategory)
	g.POST("", SaveCategory)
	g.PUT("/:id", UpdateCategory)
	g.DELETE("/:id", DeleteCategory)
}

func GetCategories(c echo.Context) error {
	var categories []models.Category

	result := database.Database.Find(&categories)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Items not found")
	}

	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category

	result := database.Database.Find(&category, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	return c.JSON(http.StatusOK, category)
}

func SaveCategory(c echo.Context) error {
	category := new(models.Category)

	err := c.Bind(category)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid body "+err.Error())
	}

	result := database.Database.Create(category)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "Invalid "+result.Error.Error())
	}

	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	result := database.Database.Find(&category, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	values := new(models.Category)
	err := c.Bind(values)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid body "+err.Error())
	}

	category.Name = values.Name
	category.Description = values.Description
	database.Database.Save(&category)

	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category

	result := database.Database.Delete(&category, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "item deleted"})
}
