package router

import (
	"task_manager/controllers"
	"task_manager/data"

	"github.com/gin-gonic/gin"
)


func SetupRouter(taskManager *data.TaskManager) *gin.Engine{
	router := gin.Default()

	handlers := controllers.NewTaskHandlers(taskManager)

	router.GET("/tasks", handlers.GetTasks)
	router.GET("/tasks/:id", handlers.GetTaskByID )
	router.PUT("tasks/:id", handlers.UpdateTask)
	router.DELETE("tasks/:id", handlers.DeleteTask)
	router.POST("/tasks", handlers.CreateTask)

	return router
}