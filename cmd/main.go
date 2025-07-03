package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"news-portal/config"
	"news-portal/internal/db"
	"news-portal/internal/delivery"
	"news-portal/internal/repository"
	"news-portal/internal/service"
)

func main() {
	// Загружаем .env конфиг
	config.LoadEnv()

	// Инициализация БД
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	// Инициализация репозиториев
	newsRepo := repository.NewNewsRepository(dbConn)

	// Инициализация сервисов
	categoryUC := service.NewCategoryService(newsRepo)
	tagUC := service.NewTagService(newsRepo)
	newsUC := service.NewNewsService(newsRepo, categoryUC, tagUC)

	// Инициализация Gin
	r := gin.Default()
	api := r.Group("/api")

	// Регистрация маршрутов
	delivery.RegisterNewsRoutes(api, newsUC, categoryUC, tagUC)
	delivery.RegisterCategoryRoutes(api, categoryUC)
	delivery.RegisterTagRoutes(api, tagUC)

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
