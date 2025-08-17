package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
	createUserController     *CreateUserController
}

func NewUserController() *UserController {
	createUserController := NewCreateUserController()

	return &UserController{
		createUserController:     createUserController,
	}
}

func (uc *UserController) SetupRoutes(router *gin.Engine) {
	router.POST("/users", uc.createUserController.Exec)
}
