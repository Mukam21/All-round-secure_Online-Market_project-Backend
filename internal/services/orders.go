package services

import (
	db "online-Market_project_Golang-Backent/internal/db"
	"online-Market_project_Golang-Backent/internal/models"
)

func GetOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := db.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// CreateOrder создает новый заказ
func CreateOrder(order models.Order) error {
	result := db.DB.Create(&order)
	return result.Error
}

// UpdateOrderStatus обновляет статус заказа
func UpdateOrderStatus(orderID uint, status string) error {
	var order models.Order
	if err := db.DB.First(&order, orderID).Error; err != nil {
		return err
	}
	order.Status = status
	result := db.DB.Save(&order)
	return result.Error
}

func GetOrderByID(id uint) (models.Order, error) {
	var order models.Order
	if err := db.DB.First(&order, id).Error; err != nil {
		return order, err
	}
	return order, nil
}

// DeleteOrder удаляет заказ по ID
func DeleteOrder(id uint) error {
	result := db.DB.Delete(&models.Order{}, id)
	return result.Error
}
