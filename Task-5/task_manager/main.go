package main

import (
	"log"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	client, err := data.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	taskManager := data.NewTaskManager(client)
	router:= router.SetupRouter(taskManager)
	router.Run("localhost:8080")
}