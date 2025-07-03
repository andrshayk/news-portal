package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv загружает переменные окружения из .env файла
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables.")
	}
}
