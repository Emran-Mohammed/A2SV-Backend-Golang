package router

import (
	"task_manager/Delivery/controllers"
	infrastructure "task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)


func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController, JWTSecret []byte) *gin.Engine{
	router := gin.Default()


	

	auth:=router.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
	}

	api := router.Group("/api")
	api.Use(infrastructure.RequireAuth(JWTSecret))
	{
		api.GET("/tasks", taskController.GetTasks)
		api.GET("/tasks/:id", taskController.GetTaskByID)

		admin := api.Group("/admin")
		admin.Use(infrastructure.RequireRole("admin"))
		{
			admin.PUT("/tasks/:id", taskController.UpdateTask)
			admin.DELETE("/tasks/:id", taskController.DeleteTask)
			admin.POST("/tasks", taskController.CreateTask)

		}

	}

	return router
}