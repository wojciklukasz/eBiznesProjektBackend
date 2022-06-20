package controllers

import (
	"ProjektBackend/api/v1/database"
	"ProjektBackend/api/v1/models"
)

func GetItemOrderByID(id int) (models.ItemOrder, error) {
	var itemOrder models.ItemOrder

	result := database.Database.First(&itemOrder, id)
	if result.Error != nil {
		return itemOrder, result.Error
	}

	return itemOrder, nil
}

func SaveItemOrder(productID int, count int) models.ItemOrder {
	itemOrder := new(models.ItemOrder)
	itemOrder.ProductID = productID
	itemOrder.Count = count
	database.Database.Create(itemOrder)

	return *itemOrder
}
