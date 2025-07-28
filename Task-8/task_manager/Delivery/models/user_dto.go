package models

import domain "task_manager/Domain"

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role	string	`json:"role"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func DTOJsonToDomainUser(userReq *UserRequest) *domain.User{
	return &domain.User{
		Username: userReq.Username,
		Password: userReq.Password,
		Role: domain.Role(userReq.Role),
	}
}