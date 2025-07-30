package mocks

import "github.com/stretchr/testify/mock"

type MockPassword struct {
	mock.Mock
}


func (m * MockPassword) HashPassword (password string)(string, error){
	args := m.Called(password)
	if args.Get(0) != "" {
		return args.Get(0).(string), args.Error(1)
	}
	return "", args.Error(1)

}

func (m *MockPassword) VerifyPassword(hashedPassword, password string) bool{
	args:= m.Called(hashedPassword, password)
	return args.Get(0).(bool)

}