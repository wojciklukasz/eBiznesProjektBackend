package controllers

import (
	"ProjektBackend/api/v1/database"
	"ProjektBackend/api/v1/models"
)

func SaveItemOrder(productID int, count int) models.ItemOrder {
	itemOrder := new(models.ItemOrder)
	itemOrder.ProductID = productID
	itemOrder.Count = count
	database.Database.Create(itemOrder)

	return *itemOrder
}
