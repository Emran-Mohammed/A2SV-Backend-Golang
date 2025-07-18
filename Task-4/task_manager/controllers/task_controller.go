package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"
	"github.com/gin-gonic/gin"
)

var taskManager = data.NewTaskManager()



func GetTasks(c *gin.Context){
	tasks := taskManager.GetTasks()
	c.IndentedJSON(http.StatusOK, tasks)

}
func GetTaskByID(c *gin.Context){
	idstr := c.Param("id")
	id, err:= strconv.Atoi(idstr)
	if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
        return
    }
	task, err2 := taskManager.GetTaskByID(id)
	if err2 != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err2.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}

func UpdateTask(c *gin.Context){
	//since the method is put i just replace execpt the id

	idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
        return
    }
	var updatedTask models.Task
	if err:= c.ShouldBindJSON(&updatedTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if !updatedTask.Status.IsValid() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid status value"})
		return
	}
	err2 := taskManager.UpdateTask(id, updatedTask)
	if err2 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err2.Error()})
		return
	}
	updatedTask.ID = id
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "upadted successfully",
		"task": updatedTask,

	})

}

func DeleteTask(c *gin.Context){
	idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
        return
    }
	taskManager.DeleteTask(id)
	c.Status(http.StatusNoContent)

}

func CreateTask(c *gin.Context){
	var task models.Task
	if err:= c.ShouldBindJSON(&task); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
		
	}
	if !task.Status.IsValid() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid status value"})
		return
	}
	id := taskManager.CreateTask(task)
	task.ID = id

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "created successfully",
	"task": task,
	})

}