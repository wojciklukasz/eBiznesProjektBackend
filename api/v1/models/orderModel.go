package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Total   float32     `gorm:"not null" json:"total"`
	Name    string      `gorm:"not null" json:"name"`
	Surname string      `gorm:"not null" json:"surname"`
	Email   string      `gorm:"not null" json:"email"`
	Nr      string      `gorm:"not null" json:"nr"`
	Road    string      `gorm:"not null" json:"road"`
	Code    string      `gorm:"not null" json:"code"`
	City    string      `gorm:"not null" json:"city"`
	Phone   string      `gorm:"not null" json:"phone"`
	Items   []ItemOrder `gorm:"foreignKey:ID"`
}
