package db

import (
	"fmt"
	"log"
	"online-Market_project_Golang-Backent/internal/config"
	"online-Market_project_Golang-Backent/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Ошибка при подключении к базе данных: %v", err)
		return err
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.CartItem{},
		&models.Order{},
		&models.Product{},
	)
	if err != nil {
		log.Printf("Ошибка при миграции базы данных: %ව", err)
		return err
	}

	log.Println("База данных успешно инициализирована")
	return nil
}
