package handlers

import (
	"net/http"
	"online-Market_project_Golang-Backent/internal/models"
	"online-Market_project_Golang-Backent/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не аутентифицирован"})
		return
	}
	u := user.(models.User)
	cart, err := services.GetCart(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить корзину"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func AddToCart(c *gin.Context) {
	var item models.CartItem
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не аутентифицирован"})
		return
	}
	u := user.(models.User)
	item.UserID = u.ID
	if err := services.AddToCart(item.UserID, uint(item.ProductID), item.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить товар в корзину"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Товар добавлен в корзину"})
}

func UpdateCartItem(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID элемента корзины"})
		return
	}
	var item models.CartItem
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	item.ID = uint(itemID)
	if err := services.UpdateCartItem(item.ID, item.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить элемент корзины"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Элемент корзины обновлен"})
}

func DeleteCartItem(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID элемента корзины"})
		return
	}
	if err := services.DeleteCartItem(uint(itemID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить элемент из корзины"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Элемент удален из корзины"})
}
