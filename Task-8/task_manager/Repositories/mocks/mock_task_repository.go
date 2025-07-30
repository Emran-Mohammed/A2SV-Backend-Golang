package mocks

import (
	"context"
	"task_manager/Domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	args := m.Called(ctx, task)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskRepository) GetTasks(ctx context.Context) ([]domain.Task, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
        return args.Get(0).([]domain.Task), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(ctx context.Context, id string, task *domain.Task) (*domain.Task, error) {
	args := m.Called(ctx, id, task)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Task), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTaskRepository) DeleteTask(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
