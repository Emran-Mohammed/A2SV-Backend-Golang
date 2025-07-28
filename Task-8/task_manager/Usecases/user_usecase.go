package usecases

import (
	"context"
	"fmt"
	"task_manager/Domain"
)




type userUsercase struct{
	userRepo domain.IUserRepository
	passwordService domain.IPasswordService
	jwtservice domain.IJWTService
}

func NewUserUsecase(repo domain.IUserRepository,pass domain.IPasswordService, jwt domain.IJWTService ) domain.IUserUsercase{
	return &userUsercase{userRepo: repo, passwordService: pass, jwtservice: jwt}
}

func (u *userUsercase) Register(ctx context.Context, user *domain.User) error{
	hashed, err := u.passwordService.HashPassword(user.Password)
	if err != nil{
		return err
	}
	user.Password = hashed
	return u.userRepo.Register(ctx, user)
}
func(u *userUsercase) Login(ctx context.Context, user *domain.User) (string, error){
	existingUser, err := u.userRepo.Login(ctx, user)
	if err != nil {
		return "", err
	}

	if !u.passwordService.VerifyPassword(existingUser.Password, user.Password){
		return "", fmt.Errorf("invalid credentials") 
	}
	token, err:= u.jwtservice.GenerateToken(user)
	if err != nil{
		return "", nil
	}
	return token, nil
}