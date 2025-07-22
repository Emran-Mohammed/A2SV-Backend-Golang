package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"
	"github.com/gin-gonic/gin"
	"time"
	"context"
)

type authHandler struct {
	userManager * data.UserManager
}

func NewUserHandler(um *data.UserManager)*authHandler{
	return &authHandler{userManager: um}
}


func (ah *authHandler) Register(c *gin.Context ){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var user models.User
	if err:= c.ShouldBindJSON(&user); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := ah.userManager.Register(ctx, user)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (ah *authHandler) Login(c *gin.Context){
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var user models.User
	if err:= c.ShouldBindJSON(&user); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	jwtToken, err := ah.userManager.Login(ctx, user)
	if err != nil{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message":"user logged in successfully","token":*jwtToken})

}