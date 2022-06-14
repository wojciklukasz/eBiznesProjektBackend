package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email   string `gorm:"not null"`
	Service string `gorm:"not null"`
	GoToken string `gorm:"not null"`
}
