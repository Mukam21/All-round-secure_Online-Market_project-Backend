package handlers

import (
	"net/http"
	"online-Market_project_Golang-Backent/internal/models"
	"online-Market_project_Golang-Backent/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не аутентифицирован"})
		return
	}
	u := user.(models.User)
	products, err := services.GetAllProducts(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить товары"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func Protected(c *gin.Context) {
	c.JSON(200, gin.H{"message": "This is a protected route"})
}

func GetProduct(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не аутентифицирован"})
		return
	}
	u := user.(models.User)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID товара"})
		return
	}
	product, err := services.GetProductByID(uint(id), u.ID)
	if err != nil {
		if err.Error() == "доступ к товару ограничен по возрасту" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ к товару ограничен по возрасту"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	if err := services.CreateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать товар"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Товар создан"})
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID товара"})
		return
	}
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	product.ID = uint(id)
	if err := services.UpdateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить товар"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Товар обновлен"})
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID товара"})
		return
	}
	if err := services.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить товар"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Товар удален"})
}
