package controllers

import (
	"net/http"
	"task_manager/Delivery/models"
	domain "task_manager/Domain"
	"github.com/gin-gonic/gin"
)
type TaskController struct{
	taskUsecase domain.ITaskUsecase
}

func NewTaskController(tuc domain.ITaskUsecase) *TaskController{
	return &TaskController{taskUsecase: tuc}
}



func (tc *TaskController) CreateTask (c *gin.Context){
	ctx := c.Request.Context()
	var createTask models.TaskRequest
	if err:= c.ShouldBindJSON(&createTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return	
	}
	
	createdDomainTask, err := tc.taskUsecase.CreateTask(ctx, models.DTOJsonToDomainTask(&createTask))
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "created successfully",
	"task": models.DTODomainToJsonTask(createdDomainTask),
	})

}
func (tc *TaskController) GetTasks(c *gin.Context){
	ctx := c.Request.Context()
	tasks, err := tc.taskUsecase.GetTasks(ctx)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)

}
func (tc *TaskController) GetTaskByID(c *gin.Context){
	ctx := c.Request.Context()
	id := c.Param("id")
	task, err := tc.taskUsecase.GetTaskByID(ctx, id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}

func (tc *TaskController) UpdateTask(c *gin.Context){
	//since the method is put i just replace execpt the id
	ctx := c.Request.Context()
	id := c.Param("id")
    
	var updatedTask models.TaskRequest
	if err:= c.ShouldBindJSON(&updatedTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	updatedDomainTask, err := tc.taskUsecase.UpdateTask(ctx, id, models.DTOJsonToDomainTask(&updatedTask))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "upadted successfully",
		"task":models.DTODomainToJsonTask(updatedDomainTask),

	})

}

func (tc *TaskController) DeleteTask(c *gin.Context){
	ctx := c.Request.Context()
	id := c.Param("id")
   
	if err := tc.taskUsecase.DeleteTask(ctx, id); err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.Status(http.StatusNoContent)

}
