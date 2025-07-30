package mocks

import (
	"context"
	domain "task_manager/Domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository)Register(ctx context.Context, user *domain.User) error{
	args := m.Called(ctx, user)
	return args.Error(0)

}

func (m * MockUserRepository)Login(ctx context.Context, user *domain.User) (*domain.User, error){
	args:= m.Called(ctx, user)
	if args.Get(0) != nil{
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}