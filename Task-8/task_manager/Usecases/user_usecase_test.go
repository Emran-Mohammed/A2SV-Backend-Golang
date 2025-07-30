package usecases_test

import (
	"context"
	"errors"
	domain "task_manager/Domain"
	infraMocks "task_manager/Infrastructure/mocks"
	repoMocks "task_manager/Repositories/mocks"
	usecases "task_manager/Usecases"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userUsecaseTestSuite struct {
	suite.Suite
	mockRepo *repoMocks.MockUserRepository
	mockPass *infraMocks.MockPassword
	mockJwt *infraMocks.MockJwtService
	usecase domain.IUserUsecase
	ctx context.Context
}

func(s *userUsecaseTestSuite) SetupTest(){
	s.ctx = context.Background()
	s.mockRepo = new(repoMocks.MockUserRepository)
	s.mockPass = new(infraMocks.MockPassword)
	s.mockJwt = new(infraMocks.MockJwtService)
	s.usecase = usecases.NewUserUsecase(s.mockRepo, s.mockPass, s.mockJwt)
}

func TestUserusecaseTestSuit(t *testing.T){
	suite.Run(t, new(userUsecaseTestSuite))
}

func (s *userUsecaseTestSuite)TestRegister_Success(){
	user := &domain.User{
		Username: "test",
		Password: "password",
		Role: domain.RoleUser,
	}
	s.mockPass.On("HashPassword", user.Password).Return("hashedpassword", nil)
	s.mockRepo.On("Register", s.ctx, mock.AnythingOfType("*domain.User")).Return(nil)

	err := s.usecase.Register(s.ctx, user)

	s.Require().NoError(err)




}
func (s *userUsecaseTestSuite)TestRegister_Failure(){
		user := &domain.User{
        Username: "test",
        Password: "password",
        Role:     domain.RoleUser,
    }

    expectedErr := errors.New("registration failed")
	    s.mockPass.On("HashPassword", user.Password).Return("hashedpassword", nil)
    s.mockRepo.On("Register", s.ctx, mock.AnythingOfType("*domain.User")).Return(expectedErr)

    err := s.usecase.Register(s.ctx, user)

    s.Require().Error(err)
    s.Equal(expectedErr, err)

}


func (s *userUsecaseTestSuite) TestLogin_Success() {
	user := &domain.User{
        Username: "test",
        Password: "password",
        Role:     domain.RoleUser,
    }
    existingUser := &domain.User{
        Username: "test",
        Password: "hashedpassword",
        Role:     domain.RoleUser,
    }

   
    s.mockRepo.On("Login", s.ctx, user).Return(existingUser, nil)
   
    s.mockPass.On("VerifyPassword", existingUser.Password, user.Password).Return(true)
   
    s.mockJwt.On("GenerateToken", user).Return("mocked_token", nil)

    token, err := s.usecase.Login(s.ctx, user)

    s.Require().NoError(err)
    s.Equal("mocked_token", token)
}

func (s *userUsecaseTestSuite) TestLogin_Failure() {
    user := &domain.User{
        Username: "test",
        Password: "password",
        Role:     domain.RoleUser,
    }

    s.mockRepo.On("Login", s.ctx, user).Return(nil, errors.New("user not found"))

    token, err := s.usecase.Login(s.ctx, user)

    s.Require().Error(err)
    s.Empty(token)
}

func (s *userUsecaseTestSuite) TestRegister_HashPasswordFails() {
    user := &domain.User{
        Username: "test",
        Password: "password",
        Role:     domain.RoleUser,
    }

    s.mockPass.On("HashPassword", user.Password).Return("", errors.New("hash failed"))

    err := s.usecase.Register(s.ctx, user)

    s.Require().Error(err)
    s.EqualError(err, "hash failed")
}

func (s *userUsecaseTestSuite) TestLogin_IncorrectPassword() {
    user := &domain.User{
        Username: "test",
        Password: "wrong_password",
    }

    existingUser := &domain.User{
        Username: "test",
        Password: "hashed_password",
    }

    s.mockRepo.On("Login", s.ctx, user).Return(existingUser, nil)
    s.mockPass.On("VerifyPassword", existingUser.Password, user.Password).Return(false)

    token, err := s.usecase.Login(s.ctx, user)

    s.Require().Error(err)
    s.Empty(token)
}
func (s *userUsecaseTestSuite) TestLogin_GenerateTokenFailed(){
	user := &domain.User{
        Username: "test",
        Password: "wrong_password",
    }

    existingUser := &domain.User{
        Username: "test",
        Password: "hashed_password",
    }

	s.mockRepo.On("Login", s.ctx, user).Return(existingUser,nil)
	s.mockPass.On("VerifyPassword",existingUser.Password, user.Password).Return(true)
	s.mockJwt.On("GenerateToken",user).Return("",errors.New("genetate token fails"))

	token, err := s.usecase.Login(s.ctx, user)

	s.Require().Error(err)
	s.Empty(token)


}
