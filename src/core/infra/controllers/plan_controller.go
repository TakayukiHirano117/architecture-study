package controllers

import "github.com/gin-gonic/gin"

type PlanController struct {
	createPlanController *CreatePlanController
	getPlansController   *GetPlansController
}

func NewPlanController() *PlanController {
	createPlanController := NewCreatePlanController()
	getPlansController := NewGetPlansController()
	return &PlanController{
		createPlanController: createPlanController,
		getPlansController:   getPlansController,
	}
}

func (pc *PlanController) SetupRoutes(router *gin.Engine) {
	router.POST("/plans", pc.createPlanController.Exec)
	router.GET("/plans", pc.getPlansController.Exec)
}
