package models

import (
	domain "task_manager/Domain"
	"time"
)

type TaskRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"duedate" binding:"required"`
	Status      string    `json:"status" binding:"required,oneof=progress completed pending"`
}

type TaskResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"duedate"`
	Status      string    `json:"status"`
}

func DTOJsonToDomainTask (taskReq *TaskRequest) *domain.Task{
	return &domain.Task{
		Title: taskReq.Title,
		Description: taskReq.Description,
		DueDate: taskReq.DueDate,
		Status: domain.Status(taskReq.Status),
	}
}

func DTODomainToJsonTask (task *domain.Task) *TaskResponse {
	return & TaskResponse{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
		DueDate: task.DueDate,
		Status: string(task.Status),
	}
}