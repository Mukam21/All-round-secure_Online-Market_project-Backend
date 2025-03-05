package services

import (
	db "online-Market_project_Golang-Backent/internal/db"
	"online-Market_project_Golang-Backent/internal/models"
)

func AddToCart(userID uint, productID uint, quantity int) error {
	var existingItem models.CartItem
	if err := db.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&existingItem).Error; err == nil {
		existingItem.Quantity += quantity
		return db.DB.Save(&existingItem).Error
	}
	cartItem := models.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}
	result := db.DB.Create(&cartItem)
	return result.Error
}

func UpdateCartItem(cartItemID uint, quantity int) error {
	var cartItem models.CartItem
	if err := db.DB.First(&cartItem, cartItemID).Error; err != nil {
		return err
	}
	cartItem.Quantity = quantity
	result := db.DB.Save(&cartItem)
	return result.Error
}

func GetCart(userID uint) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	if err := db.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}

func DeleteCartItem(cartItemID uint) error {
	var cartItem models.CartItem
	if err := db.DB.First(&cartItem, cartItemID).Error; err != nil {
		return err
	}
	result := db.DB.Delete(&cartItem)
	return result.Error
}
