package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"task_manager/Domain"
	"task_manager/Repositories/mocks"
	"task_manager/Usecases"
)

type taskUsecaseTestSuite struct {
	suite.Suite
	mockRepo *mocks.MockTaskRepository
	usecase domain.ITaskUsecase
	ctx context.Context
}
func (s *taskUsecaseTestSuite) SetupTest(){
	s.ctx = context.Background()
	s.mockRepo = new(mocks.MockTaskRepository)
	s.usecase = usecases.NewTaskUsecase(s.mockRepo)

}
func TestTaskUsecaseTestSuite(t *testing.T){
	suite.Run(t, new(taskUsecaseTestSuite))
}
func (s *taskUsecaseTestSuite) TestCreateTask_Success(){
	task := &domain.Task{
		Title: "Test Task",
		Description: "This is a test",
		DueDate: time.Now().Add(24 * time.Hour),
		Status: domain.Pending,
	}
	s.mockRepo.On("CreateTask", s.ctx, mock.AnythingOfType("*domain.Task")).Return(task, nil)

	result, err := s.usecase.CreateTask(s.ctx, task)
	s.Require().NoError(err)
	s.Require().Equal(task.Title, result.Title)
	s.Require().Equal(task.Description, result.Description)
	s.Require().Equal(task.DueDate, result.DueDate)
	s.Require().Equal(task.Status, result.Status)

}
func (s *taskUsecaseTestSuite) TestCreateTask_Failure() {
	task := &domain.Task{
		Title: "Failing Task",
	}

	s.mockRepo.On("CreateTask", s.ctx, mock.AnythingOfType("*domain.Task")).
		Return(nil, errors.New("DB error"))

	result, err := s.usecase.CreateTask(s.ctx, task)
	s.Error(err)
	s.Nil(result)
	s.Contains(err.Error(), "DB error")
}
func (s *taskUsecaseTestSuite) TestGetTasks_Success() {
	tasks := []domain.Task{
		{ID: "1", Title: "Task 1"},
		{ID: "2", Title: "Task 2"},
	}
	s.mockRepo.On("GetTasks", s.ctx).Return(tasks, nil)

	result, err := s.usecase.GetTasks(s.ctx)
	s.Require().NoError(err)
	s.Equal(2, len(result))
	s.mockRepo.AssertExpectations(s.T())
}
func (s *taskUsecaseTestSuite) TestGetTasks_Failure() {
	s.mockRepo.On("GetTasks", s.ctx).Return(nil, errors.New("DB failed"))

	result, err := s.usecase.GetTasks(s.ctx)
	s.Error(err)
	s.Nil(result)
}

func (s *taskUsecaseTestSuite) TestGetTaskByID_Success() {
	task := &domain.Task{ID: "abc123", Title: "My Task"}
	s.mockRepo.On("GetTaskByID", s.ctx, "abc123").Return(task, nil)

	result, err := s.usecase.GetTaskByID(s.ctx, "abc123")
	s.NoError(err)
	s.Equal(task.ID, result.ID)
}
func (s *taskUsecaseTestSuite) TestGetTaskByID_NotFound() {
	s.mockRepo.On("GetTaskByID", s.ctx, "invalid-id").Return(nil, errors.New("task not found"))

	result, err := s.usecase.GetTaskByID(s.ctx, "invalid-id")
	s.Error(err)
	s.Nil(result)
	s.Contains(err.Error(), "task not found")
}

func (s *taskUsecaseTestSuite) TestUpdateTask_Success() {
	task := &domain.Task{ID: "abc123", Title: "Updated"}
	s.mockRepo.On("UpdateTask", s.ctx, "abc123", task).Return(task, nil)

	result, err := s.usecase.UpdateTask(s.ctx, "abc123", task)
	s.NoError(err)
	s.Equal("Updated", result.Title)
}
func (s *taskUsecaseTestSuite) TestUpdateTask_Failure() {
	task := &domain.Task{Title: "Broken update"}
	s.mockRepo.On("UpdateTask", s.ctx, "invalid-id", task).Return(nil, errors.New("update failed"))

	result, err := s.usecase.UpdateTask(s.ctx, "invalid-id", task)
	s.Error(err)
	s.Nil(result)
}

func (s *taskUsecaseTestSuite) TestDeleteTask_Success() {
	s.mockRepo.On("DeleteTask", s.ctx, "abc123").Return(nil)

	err := s.usecase.DeleteTask(s.ctx, "abc123")
	s.NoError(err)
}

func (s *taskUsecaseTestSuite) TestDeleteTask_Failure() {
	s.mockRepo.On("DeleteTask", s.ctx, "invalid-id").Return(errors.New("delete failed"))

	err := s.usecase.DeleteTask(s.ctx, "invalid-id")
	s.Error(err)
	s.Contains(err.Error(), "delete failed")
}