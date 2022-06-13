package controllers

import (
	"ProjektBackend/api/v1/database"
	"ProjektBackend/api/v1/models"
	"github.com/google/uuid"
)

func FindUser(email string, service string) bool {
	var user models.User
	database.Database.Find(&user, "Email = ? AND Service = ?", email, service)
	if user.Email == "" {
		return false
	}
	return true
}

func GetUser(email string, service string) models.User {
	var user models.User
	database.Database.Find(&user, "Email = ? AND Service = ?", email, service)
	return user
}

func AddUser(email string, service string) models.User {
	user := new(models.User)
	user.Email = email
	user.Service = service
	user.GoToken = uuid.NewString()
	database.Database.Create(user)
	return GetUser(email, service)
}
