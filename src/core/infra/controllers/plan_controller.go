package controllers

import "github.com/gin-gonic/gin"

type PlanController struct {
	createPlanController *CreatePlanController
}

func NewPlanController() *PlanController {
	createPlanController := NewCreatePlanController()
	return &PlanController{
		createPlanController: createPlanController,
	}
}

func (pc *PlanController) SetupRoutes(router *gin.Engine) {
	router.POST("/plans", pc.createPlanController.Exec)
}
