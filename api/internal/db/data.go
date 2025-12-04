package db

import (
	"SportsClubs/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	log.Println("Подключение к базе данных прошло успешно.")

	if err := database.AutoMigrate(&models.Club{}, &models.Town{}); err != nil {
		log.Fatalf("Не удалось выполнить миграцию: %v", err)
	}
	log.Println("Миграция базы данных успешно завершена")

	return database
}
