package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	healthController *HealthController
	userController   *UserController
}

func NewController() *Controller {
	healthController := NewHealthController()
	userController := NewUserController()

	return &Controller{
		healthController: healthController,
		userController:   userController,
	}
}

func (c *Controller) SetupRoutes(router *gin.Engine) {
	router.GET("/health", c.healthController.HealthCheck)
	c.userController.SetupRoutes(router)
}
