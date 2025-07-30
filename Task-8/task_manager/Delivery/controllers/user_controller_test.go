package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/models"
	domain "task_manager/Domain"
	mockUsecase "task_manager/Usecases/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	mockUserUsecase *mockUsecase.IUserUsecase
	controller      *controllers.UserController
	router          *gin.Engine
}

func (s *UserControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.mockUserUsecase = new(mockUsecase.IUserUsecase)
	s.controller = controllers.NewUserController(s.mockUserUsecase)

	r := gin.Default()
	r.POST("/auth/register", s.controller.Register)
	r.POST("/auth/login", s.controller.Login)
	s.router = r
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (s *UserControllerTestSuite) Test_Register_Success() {
	reqBody := models.UserRequest{
		Username: "testuser",
		Password: "testpass",
		Role:     "user",
	}

	s.mockUserUsecase.On("Register", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
		return u.Username == reqBody.Username && u.Password == reqBody.Password && string(u.Role) == reqBody.Role
	})).Return(nil)

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusCreated, rec.Code)
	s.Contains(rec.Body.String(), "User registered successfully")
	s.mockUserUsecase.AssertExpectations(s.T())
}

func (s *UserControllerTestSuite) Test_Register_Failure() {
	reqBody := models.UserRequest{
		Username: "testuser",
		Password: "testpass",
		Role:     "user",
	}

	s.mockUserUsecase.On("Register", mock.Anything, mock.Anything).Return(errors.New("registration failed"))

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusBadRequest, rec.Code)
	s.Contains(rec.Body.String(), "registration failed")
	s.mockUserUsecase.AssertExpectations(s.T())
}

func (s *UserControllerTestSuite) Test_Login_Success() {
	reqBody := models.UserRequest{
		Username: "testuser",
		Password: "testpass",
		Role:     "user",
	}

	s.mockUserUsecase.On("Login", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
		return u.Username == reqBody.Username && u.Password == reqBody.Password
	})).Return("mocked_jwt_token", nil)

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	s.Contains(rec.Body.String(), "mocked_jwt_token")
	s.mockUserUsecase.AssertExpectations(s.T())
}

func (s *UserControllerTestSuite) Test_Login_Failure() {
	reqBody := models.UserRequest{
		Username: "wronguser",
		Password: "wrongpass",
		Role:     "user",
	}

	s.mockUserUsecase.On("Login", mock.Anything, mock.Anything).Return("", errors.New("login faild"))

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusUnauthorized, rec.Code)
	s.Contains(rec.Body.String(), "login faild")
	s.mockUserUsecase.AssertExpectations(s.T())
}
