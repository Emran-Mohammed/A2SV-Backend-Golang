package usecases

import (
	"context"
	"task_manager/Domain"
)




type taskUsecase struct {
	taskRepo domain.ITaskRepository
}

func NewTaskUsecase(repo domain.ITaskRepository) domain.ITaskUsecase {
	return &taskUsecase{taskRepo: repo}
}

func (u *taskUsecase) CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error){
	return u.taskRepo.CreateTask(ctx, task)
}
func (u *taskUsecase) GetTasks(ctx context.Context) ([]domain.Task, error) {
	return u.taskRepo.GetTasks(ctx)
}

func (u *taskUsecase) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	return u.taskRepo.GetTaskByID(ctx, id)
}

func (u *taskUsecase) UpdateTask(ctx context.Context, id string, task *domain.Task) (*domain.Task, error) {
	return u.taskRepo.UpdateTask(ctx, id, task)
}

func (u *taskUsecase) DeleteTask(ctx context.Context, id string) error {
	return u.taskRepo.DeleteTask(ctx, id)
}
