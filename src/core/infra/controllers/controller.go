package controllers

import (
	"net/http"
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

func (c *Controller) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", c.healthController.HealthCheck)
}
