package handlers

import (
	"net/http"
	"online-Market_project_Golang-Backent/internal/models"
	"online-Market_project_Golang-Backent/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не аутентифицирован"})
		return
	}
	u := user.(models.User)
	orders, err := services.GetOrders(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить заказы"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не аутентифицирован"})
		return
	}
	u := user.(models.User)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID заказа"})
		return
	}
	order, err := services.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}
	if order.UserID != u.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Доступ к заказу запрещён"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	if err := services.CreateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать заказ"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Заказ создан"})
}

func UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID заказа"})
		return
	}
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	order.ID = uint(id)
	if err := services.UpdateOrderStatus(order.ID, order.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить заказ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Заказ обновлен"})
}

func DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID заказа"})
		return
	}
	if err := services.DeleteOrder(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось отменить заказ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Заказ отменен"})
}
