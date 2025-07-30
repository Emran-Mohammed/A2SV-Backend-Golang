package repositories_test

import (
	"context"
	"testing"
	"time"

	"task_manager/Domain"
	"task_manager/Repositories"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type TaskRepositoryIntegrationSuite struct {
	suite.Suite
	db     *mongo.Database
	repo   domain.ITaskRepository
	ctx    context.Context
	cancel context.CancelFunc
}

func (s *TaskRepositoryIntegrationSuite) SetupSuite() {
	db, err := repositories.ConnectDB()
	s.Require().NoError(err)
	s.db = db
	s.repo = repositories.NewTaskRepository(db)
}

func (s *TaskRepositoryIntegrationSuite) SetupTest() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// Clean before each test
	s.db.Collection("taskstest").DeleteMany(s.ctx, bson.M{})
}
func (s *TaskRepositoryIntegrationSuite) TearDownTest() {
	s.cancel()
}
func TestTaskRepositoryIntegrationSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryIntegrationSuite))
}
func (s *TaskRepositoryIntegrationSuite) TestCreateAndGetTaskByID() {
	task := &domain.Task{
		Title:       "Integration Test Task",
		Description: "Testing MongoDB Insert & Get",
		Status:      domain.Pending,
	}

	createdTask, err := s.repo.CreateTask(s.ctx, task)
	s.Require().NoError(err)
	s.Require().NotNil(createdTask)
	s.NotEmpty(createdTask.ID)

	foundTask, err := s.repo.GetTaskByID(s.ctx, createdTask.ID)
	s.Require().NoError(err)
	s.Equal(createdTask.Title, foundTask.Title)
	s.Equal(createdTask.Description, foundTask.Description)
}

func (s *TaskRepositoryIntegrationSuite) TestUpdateTask() {
	task := &domain.Task{
		Title:       "Initial Title",
		Description: "Initial Desc",
		Status:      domain.Progress,
	}

	createdTask, err := s.repo.CreateTask(s.ctx, task)
	s.Require().NoError(err)

	createdTask.Title = "Updated Title"
	updated, err := s.repo.UpdateTask(s.ctx, createdTask.ID, createdTask)
	s.Require().NoError(err)
	s.Equal("Updated Title", updated.Title)
}

func (s *TaskRepositoryIntegrationSuite) TestDeleteTask() {
	task := &domain.Task{
		Title:       "Delete Me",
		Description: "Should be removed",
		Status:      domain.Completed,
	}
	createdTask, err := s.repo.CreateTask(s.ctx, task)
	s.Require().NoError(err)

	err = s.repo.DeleteTask(s.ctx, createdTask.ID)
	s.Require().NoError(err)

	_, err = s.repo.GetTaskByID(s.ctx, createdTask.ID)
	s.Error(err)
}

func (s *TaskRepositoryIntegrationSuite) TestGetTasks() {
	task1 := &domain.Task{
		Title:       "Task 1",
		Description: "First Task",
		Status:      domain.Pending,
	}
	task2 := &domain.Task{
		Title:       "Task 2",
		Description: "Second Task",
		Status:      domain.Completed,
	}

	_, err := s.repo.CreateTask(context.Background(), task1)
	s.Require().NoError(err)

	_, err = s.repo.CreateTask(context.Background(), task2)
	s.Require().NoError(err)

	tasks, err := s.repo.GetTasks(context.Background())

	s.Require().NoError(err)
	s.Require().GreaterOrEqual(len(tasks), 2)

	var foundTask1, foundTask2 bool
	for _, t := range tasks {
		if t.Title == "Task 1" && t.Description == "First Task" {
			foundTask1 = true
		}
		if t.Title == "Task 2" && t.Description == "Second Task" {
			foundTask2 = true
		}
	}

	s.True(foundTask1, "Task 1 should be in the result")
	s.True(foundTask2, "Task 2 should be in the result")
}

