package mocks

import (
	domain "task_manager/Domain"

	"github.com/stretchr/testify/mock"
)

type MockJwtService struct {
	mock.Mock
}

func (m *MockJwtService) GenerateToken(user *domain.User) (string, error){
	args := m.Called(user)
	if args.Get(0) != ""{
		return args.Get(0).(string), args.Error(1)
	}
	return "", args.Error(1)
}