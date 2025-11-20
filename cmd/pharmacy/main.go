package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/config"
)

func main() {
	db := config.SetUpDatabaseConnection()

	fmt.Println(db) // временно

	// Выполняем миграции моделей

	router := gin.Default()

	// регистрация роутов

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}

