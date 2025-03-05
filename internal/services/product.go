package services

import (
	"errors"
	db "online-Market_project_Golang-Backent/internal/db"
	"online-Market_project_Golang-Backent/internal/models"
)

func GetAllProducts(userID uint) ([]models.Product, error) {
	if db.DB == nil {
		return nil, errors.New("база данных не инициализирована")
	}
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	var products []models.Product
	query := db.DB
	if user.Age < 18 {
		query = query.Where("age_restricted = ?", false)
	}
	result := query.Find(&products)
	return products, result.Error
}

func CreateProduct(product models.Product) error {
	if db.DB == nil {
		return errors.New("база данных не инициализирована")
	}
	result := db.DB.Create(&product)
	return result.Error
}

func UpdateProduct(product models.Product) error {
	if db.DB == nil {
		return errors.New("база данных не инициализирована")
	}
	result := db.DB.Save(&product)
	return result.Error
}

func GetProductByID(id uint, userID uint) (models.Product, error) {
	if db.DB == nil {
		return models.Product{}, errors.New("база данных не инициализирована")
	}
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return models.Product{}, err
	}

	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return product, err
	}
	if user.Age < 18 && product.AgeRestricted {
		return models.Product{}, errors.New("доступ к товару ограничен по возрасту")
	}
	return product, nil
}

func DeleteProduct(id uint) error {
	if db.DB == nil {
		return errors.New("база данных не инициализирована")
	}
	result := db.DB.Delete(&models.Product{}, id)
	return result.Error
}
