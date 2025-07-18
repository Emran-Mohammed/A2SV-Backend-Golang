package data

import (
	"fmt"
	"task_manager/models"
)
func NewTaskManager() *TaskManager{
	return &TaskManager{Tasks: make(map[int]models.Task)}
}

type TaskManager struct{
	Tasks map[int]models.Task
	nextid int

}

func (t *TaskManager) GetTasks() []models.Task{
	taskList:= make([]models.Task, 0, len(t.Tasks))

	for _ , task := range t.Tasks{
		taskList = append(taskList, task)
	}
	
	return taskList

	
}

func (t* TaskManager) GetTaskByID(id int)(models.Task, error){
	task, exist := t.Tasks[id]
	if !exist{
		return task, fmt.Errorf("a task with an %v id is not found", id)
	}
	return task, nil
}

func (t * TaskManager) UpdateTask(id int, updatedTask models.Task)error{
	_ , exist := t.Tasks[id]
	if !exist{
		return fmt.Errorf("task with an id %v is not found", id)

	}
	updatedTask.ID = id
	t.Tasks[id] = updatedTask
	return nil
}

func (t *TaskManager) DeleteTask (id int){
	delete(t.Tasks, id)
}
func (t *TaskManager) CreateTask(task models.Task) int {
	t.nextid ++
	task.ID = t.nextid
	t.Tasks[t.nextid] = task
	return task.ID


}