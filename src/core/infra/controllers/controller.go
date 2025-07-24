package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	healthController *HealthController
}

func NewController() *Controller {
	healthController := NewHealthController()

	return &Controller{
		healthController: healthController,
	}
}

func (c *Controller) SetupRoutes(router *gin.Engine) {
	router.GET("/health", c.healthController.HealthCheck)
}
