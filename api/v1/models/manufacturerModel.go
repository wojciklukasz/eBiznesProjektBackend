package models

import "gorm.io/gorm"

type Manufacturer struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}
