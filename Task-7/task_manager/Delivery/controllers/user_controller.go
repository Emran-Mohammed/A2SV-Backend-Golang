package controllers

import (
	"net/http"
	"task_manager/Delivery/models"
	domain "task_manager/Domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsercase  domain.IUserUsercase
	
}

func NewUserController(uu domain.IUserUsercase)*UserController{
	return &UserController{userUsercase: uu}
}


func (uc *UserController) Register(c *gin.Context ){
	ctx := c.Request.Context()
	var user models.UserRequest
	if err:= c.ShouldBindJSON(&user); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.userUsercase.Register(ctx, models.DTOJsonToDomainUser(&user))
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uc *UserController)Login(c *gin.Context){
	ctx := c.Request.Context()

	var user models.UserRequest
	if err:= c.ShouldBindJSON(&user); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	jwtToken, err := uc.userUsercase.Login(ctx, models.DTOJsonToDomainUser(&user))
	if err != nil{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message":"user logged in successfully","token":jwtToken})

}