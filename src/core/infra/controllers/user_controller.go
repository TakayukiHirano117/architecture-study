package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
	createUserController     *CreateUserController
	findByNameUserController *FindByNameUserController
}

func NewUserController() *UserController {
	createUserController := NewCreateUserController()
	findByNameUserController := NewFindByNameUserController()

	return &UserController{
		createUserController:     createUserController,
		findByNameUserController: findByNameUserController,
	}
}

func (uc *UserController) SetupRoutes(router *gin.Engine) {
	router.POST("/users", uc.createUserController.Exec)
	router.GET("/users/:name", uc.findByNameUserController.Exec)
}
