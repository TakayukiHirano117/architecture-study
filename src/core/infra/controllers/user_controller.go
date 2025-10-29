package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
	createUserController     *CreateUserController
	updateUserController     *UpdateUserController
}

func NewUserController() *UserController {
	createUserController := NewCreateUserController()
	updateUserController := NewUpdateUserController()

	return &UserController{
		createUserController:     createUserController,
		updateUserController:     updateUserController,
	}
}

func (uc *UserController) SetupRoutes(router *gin.Engine) {
	router.POST("/users", uc.createUserController.Exec)
	router.PUT("/users/:id", uc.updateUserController.Exec)
}
