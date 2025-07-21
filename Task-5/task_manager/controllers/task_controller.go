package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"
	"github.com/gin-gonic/gin"
	"time"
	"context"
)
type TaskHandlers struct{
	taskManager *data.TaskManager

}

func NewTaskHandlers(tm *data.TaskManager) *TaskHandlers{
	return &TaskHandlers{taskManager: tm}
}



func (h *TaskHandlers) GetTasks(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	tasks, err := h.taskManager.GetTasks(ctx)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)

}
func (h *TaskHandlers) GetTaskByID(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
	task, err := h.taskManager.GetTaskByID(ctx, id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}

func (h *TaskHandlers) UpdateTask(c *gin.Context){
	//since the method is put i just replace execpt the id
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
    
	var updatedTask models.Task
	if err:= c.ShouldBindJSON(&updatedTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if !updatedTask.Status.IsValid() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid status value"})
		return
	}
	_ , err := h.taskManager.UpdateTask(ctx, id, updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "upadted successfully",

	})

}

func (h *TaskHandlers) DeleteTask(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")
   
	if err := h.taskManager.DeleteTask(ctx, id); err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.Status(http.StatusNoContent)

}

func (h *TaskHandlers) CreateTask(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var task models.Task
	if err:= c.ShouldBindJSON(&task); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
		
	}
	if !task.Status.IsValid() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid status value"})
		return
	}
	result,  err := h.taskManager.CreateTask(ctx, task)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "created successfully",
	"task id": result.InsertedID,
	})

}