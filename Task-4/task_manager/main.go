package main

import "task_manager/router"

func main() {
	router:= router.SetupRouter()
	router.Run("localhost:8080")
}