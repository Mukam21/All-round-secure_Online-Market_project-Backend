package handlers

import (
	"net/http"
	"online-Market_project_Golang-Backent/internal/models"
	"online-Market_project_Golang-Backent/internal/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := services.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось зарегистрировать пользователя"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь зарегистрирован",
		"token":   token,
	})
}
