package main

import (
	"log"
	"user/database"
	"user/handler"
	"user/storage"
	"user/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}
	defer sqlDB.Close()

	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userStorage := storage.NewUserStorage(db)
		userUsecase := usecase.NewUserUsecase(userStorage)
		userHandler := handler.NewUserHandler(userUsecase)
		userGroup.POST("/create", userHandler.CreateUserHandler)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
