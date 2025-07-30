package controllers_test

import (
	"bytes"
	// "context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/models"
	domain "task_manager/Domain"

	"task_manager/Usecases/mocks"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskControllerTestSuite struct {
	suite.Suite
	mockTaskUsecase *mocks.ITaskUsecase
	controller      *controllers.TaskController
	router          *gin.Engine
}
func (suite *TaskControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockTaskUsecase = new(mocks.ITaskUsecase)
	suite.controller = controllers.NewTaskController(suite.mockTaskUsecase)

	r := gin.Default()
	r.POST("/admin/tasks", suite.controller.CreateTask) // Adjust to match real route
	api := r.Group("/api")
	api.POST("/admin/tasks", suite.controller.CreateTask)
	api.GET("/tasks/:id", suite.controller.GetTaskByID)
	api.GET("/tasks", suite.controller.GetTasks)
	api.PUT("/admin/tasks/:id", suite.controller.UpdateTask)
	api.DELETE("/admin/tasks/:id", suite.controller.DeleteTask)
	suite.router = r
}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}

func (suite *TaskControllerTestSuite) TestCreateTask_Success() {
	input := models.TaskRequest{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate: time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}
	domainTask := models.DTOJsonToDomainTask(&input)

	// suite.mockTaskUsecase.On("CreateTask", mock.Anything, domainTask).
	// 	Return(domainTask, nil).Once()
	suite.mockTaskUsecase.On(
    "CreateTask",
    mock.Anything, // context
    mock.MatchedBy(func(task *domain.Task) bool {
        return task.Title == input.Title &&
               task.Description == input.Description &&
               task.Status == domain.Status(input.Status)
        // Optionally check DueDate with .Equal or .Truncate(time.Second)
    }),
).Return(domainTask, nil).Once()

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/api/admin/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	suite.Equal(http.StatusCreated, rec.Code)
	suite.Contains(rec.Body.String(), "created successfully")
}


func (s *TaskControllerTestSuite) Test_GetTaskByID_Success() {
    taskID := "123"

    expected := &domain.Task{
        ID:          taskID,
        Title:       "Test Task",
        Description: "Task description",
    }

    s.mockTaskUsecase.On("GetTaskByID", mock.Anything, taskID).Return(expected, nil)

    req, _ := http.NewRequest(http.MethodGet, "/api/tasks/"+taskID, nil)
    rec := httptest.NewRecorder()

    s.router.ServeHTTP(rec, req)

    s.Equal(http.StatusOK, rec.Code)

    var response models.TaskResponse
    err := json.Unmarshal(rec.Body.Bytes(), &response)
    s.NoError(err)

    s.Equal(expected.ID, response.ID)
    s.Equal(expected.Title, response.Title)
    s.Equal(expected.Description, response.Description)

    s.mockTaskUsecase.AssertExpectations(s.T())
}
func (s *TaskControllerTestSuite) Test_UpdateTask_Success() {
    taskID := "123"
    updateBody := &models.TaskRequest{
        Title:       "Updated Title",
        Description: "Updated Description",
		Status:  "pending",
    	DueDate:     time.Now().Add(24 * time.Hour),
    }

    updatedTask := &domain.Task{
        ID:          taskID,
        Title:       updateBody.Title,
        Description: updateBody.Description,
		Status:      domain.Status(updateBody.Status),
        DueDate:     updateBody.DueDate,
		
    }

    jsonBody, _ := json.Marshal(updateBody)
    req, _ := http.NewRequest(http.MethodPut, "/api/admin/tasks/"+taskID, bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    rec := httptest.NewRecorder()

    // s.mockTaskUsecase.On("UpdateTask", mock.Anything, taskID, models.DTOJsonToDomainTask(updateBody)).Return(updatedTask, nil)
	 s.mockTaskUsecase.On("UpdateTask", mock.Anything, taskID, mock.AnythingOfType("*domain.Task")).Return(updatedTask, nil)

    s.router.ServeHTTP(rec, req)

    s.Equal(http.StatusOK, rec.Code)

	type updateTaskResponse struct {
    Message string              `json:"message"`
    Task    models.TaskResponse `json:"task"`
}

	var response updateTaskResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	s.NoError(err)

	s.Equal(updatedTask.ID, response.Task.ID)
	s.Equal(updatedTask.Title, response.Task.Title)
	s.Equal(updatedTask.Description, response.Task.Description)


    s.mockTaskUsecase.AssertExpectations(s.T())
}
func (s *TaskControllerTestSuite) Test_DeleteTask_Success() {
    taskID := "123"

    s.mockTaskUsecase.On("DeleteTask", mock.Anything, taskID).Return(nil)

    req, _ := http.NewRequest(http.MethodDelete, "/api/admin/tasks/"+taskID, nil)
    rec := httptest.NewRecorder()

    s.router.ServeHTTP(rec, req)

    s.Equal(http.StatusNoContent, rec.Code)
    s.Empty(rec.Body.String())

    s.mockTaskUsecase.AssertExpectations(s.T())
}


