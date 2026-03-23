package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
	createUserController *CreateUserController
	updateUserController *UpdateUserController
}

func NewUserController() *UserController {
	createUserController := NewCreateUserController()
	updateUserController := NewUpdateUserController()

	return &UserController{
		createUserController: createUserController,
		updateUserController: updateUserController,
	}
}

// SetupPublicRoutes は認証不要のルートを登録します
func (uc *UserController) SetupPublicRoutes(router gin.IRouter) {
	router.POST("/users", uc.createUserController.Exec)
}

// SetupProtectedRoutes は認証必須のルートを登録します
func (uc *UserController) SetupProtectedRoutes(router gin.IRouter) {
	router.PUT("/users/:id", uc.updateUserController.Exec)
}
