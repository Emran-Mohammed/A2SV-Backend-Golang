package infrastructure

import (

	"golang.org/x/crypto/bcrypt"
)


type PasswordService interface{
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, password string) bool
}

type passwordService struct{}

func NewPasswordService() PasswordService{
	return &passwordService{}
}

func (p *passwordService) HashPassword(password string) (string, error){

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

func (p *passwordService) VerifyPassword(hashedPassword, password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	
	return err == nil
}