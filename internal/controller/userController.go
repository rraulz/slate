package controller

import (
	"log"
	"slate/internal/errors"
	"slate/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewUserController(userService *service.UserService, authService *service.AuthService) *UserController {
	return &UserController{
		userService: userService,
		authService: authService,
	}
}

func (uc *UserController) LoginUserController(g *gin.Context) {
	loginData := &service.LoginData{
		Username: g.PostForm("username"),
		Password: g.PostForm("password"),
	}

	_, err := uc.authService.LoginUser(loginData)
	if err != nil {
		apiError := errors.ConvertErrorToGinResponse(err)
		log.Printf("Error: %v", apiError.InternalMessage)
		// g.JSON(apiError.Status, gin.H{"error": apiError.ResponseMessage})
	}

	// g.JSON(200, gin.H{token: token})
	g.JSON(200, gin.H{"user": "user", "password": "a", "errorMessage": "a"})
}
