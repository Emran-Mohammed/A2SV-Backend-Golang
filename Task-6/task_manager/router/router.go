package router

import (
	"task_manager/config"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)


func SetupRouter(taskManager *data.TaskManager, userManager *data.UserManager) *gin.Engine{
	router := gin.Default()


	authHandler := controllers.NewUserHandler(userManager)

	taskHandlers := controllers.NewTaskHandlers(taskManager)

	

	auth:=router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login",authHandler.Login)
	}

	api := router.Group("/api")
	api.Use(middleware.RequireAuth(config.JWTSecret))
	{
		api.GET("/tasks", taskHandlers.GetTasks)
		api.GET("/tasks/:id", taskHandlers.GetTaskByID )

		admin := api.Group("/admin")
		admin.Use(middleware.RequireRole("admin"))
		{
			admin.PUT("tasks/:id", taskHandlers.UpdateTask)
			admin.DELETE("tasks/:id", taskHandlers.DeleteTask)
			admin.POST("/tasks", taskHandlers.CreateTask)

		}

	}

	return router
}