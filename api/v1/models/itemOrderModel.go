package models

import "gorm.io/gorm"

type ItemOrder struct {
	gorm.Model
	ProductID int `gorm:"not null" json:"product_id"`
	Count     int `gorm:"not null" json:"count"`
}
