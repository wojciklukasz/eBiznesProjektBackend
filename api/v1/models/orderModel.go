package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Date       time.Time `gorm:"not null" json:"date"`
	CustomerID int       `gorm:"not null" json:"customerID"`
	Total      float32   `gorm:"not null" json:"total"`
}
