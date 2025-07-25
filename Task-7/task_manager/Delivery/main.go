package main

import (
	"log"
	"os"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/router"
	infrastructure "task_manager/Infrastructure"
	repositories "task_manager/Repositories"
	usecases "task_manager/Usecases"

	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database, err := repositories.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	JWTSecret := []byte(os.Getenv("JWT_SECRET"))
	jwtService := infrastructure.NewJWTService(JWTSecret)
	passwordService := infrastructure.NewPasswordService()
	userRepository := repositories.NewUserRepository(database)
	userUsecase := usecases.NewUserUsecase(userRepository, passwordService, jwtService)
	userController := controllers.NewUserController(userUsecase)

	taskRepository := repositories.NewTaskRepository(database)
	taskUsecase := usecases.NewTaskUsecase(taskRepository)
	taskController := controllers.NewTaskController(taskUsecase)
	router:= router.SetupRouter(taskController, userController, JWTSecret)
	router.Run("localhost:8080")
}