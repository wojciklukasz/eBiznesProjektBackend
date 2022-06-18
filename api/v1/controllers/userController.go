package controllers

import (
	"ProjektBackend/api/v1/database"
	"ProjektBackend/api/v1/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUsersRouting(e *echo.Group) {
	g := e.Group("/user")
	g.DELETE("/:id", DeleteUser)
	g.GET("/validate/:token", FindUserWithToken)
}

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

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	result := database.Database.Delete(&user, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, itemNotFoundMessage)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "item deleted"})
}

func FindUserWithToken(c echo.Context) error {
	token := c.Param("token")
	var user models.User

	AddUser("a@b.com", "custom")

	database.Database.Find(&user, "go_token = ?", token)
	if user.Email == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "invalid token"})
	} else {
		return c.JSON(http.StatusOK, map[string]string{"email": user.Email})
	}
}
