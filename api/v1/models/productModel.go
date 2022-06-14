package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name           string  `gorm:"not null" json:"name"`
	Price          float32 `gorm:"not null" json:"price"`
	CategoryID     int     `gorm:"not null" json:"categoryID"`
	ManufacturerID int     `gorm:"not null" json:"manufacturerID"`
	Description    string  `gorm:"not null" json:"description"`
}
