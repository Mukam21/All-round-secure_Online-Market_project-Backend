package main

import (
	"log"
	"online-Market_project_Golang-Backent/internal/config"
	"online-Market_project_Golang-Backent/internal/db"
	"online-Market_project_Golang-Backent/internal/handlers"
	"online-Market_project_Golang-Backent/internal/middleware"
	"online-Market_project_Golang-Backent/internal/parser"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("./cmd/config.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация базы данных
	if err := db.InitDatabase(cfg); err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	// Запуск парсера для двух категорий (можно закомментировать после первого запуска)
	parser.ParseLenta("https://lenta.com/catalog/product/")
	parser.ParseLenta("https://lenta.com/catalog/products-prod/")

	// Инициализация Gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies([]string{"127.0.0.1"}) // Указываем доверенные прокси

	r.POST("/register", handlers.Register)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		protected.GET("/protected", handlers.Protected)

		protected.GET("/products", handlers.GetProducts)
		protected.GET("/products/:id", handlers.GetProduct)
		protected.POST("/products", handlers.CreateProduct)
		protected.PUT("/products/:id", handlers.UpdateProduct)
		protected.DELETE("/products/:id", handlers.DeleteProduct)

		protected.GET("/cart", handlers.GetCart)
		protected.POST("/cart", handlers.AddToCart)
		protected.PUT("/cart/:item_id", handlers.UpdateCartItem)
		protected.DELETE("/cart/:item_id", handlers.DeleteCartItem)

		protected.GET("/orders", handlers.GetOrders)
		protected.GET("/orders/:id", handlers.GetOrder)
		protected.POST("/orders", handlers.CreateOrder)
		protected.PUT("/orders/:id", handlers.UpdateOrder)
		protected.DELETE("/orders/:id", handlers.DeleteOrder)
	}

	// Запуск сервера
	log.Printf("Запуск сервера на порту %s", cfg.ServerPort)
	if err := r.Run(cfg.ServerPort); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
