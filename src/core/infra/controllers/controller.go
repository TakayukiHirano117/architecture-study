package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	healthController            *HealthController
	userController              *UserController
	mentorRecruitmentController *MentorRecruitmentController
	planController              *PlanController
}

func NewController() *Controller {
	healthController := NewHealthController()
	userController := NewUserController()
	mentorRecruitmentController := NewMentorRecruitmentController()
	planController := NewPlanController()

	return &Controller{
		healthController:            healthController,
		userController:              userController,
		mentorRecruitmentController: mentorRecruitmentController,
		planController:              planController,
	}
}

func (c *Controller) SetupRoutes(router *gin.Engine) {
	router.GET("/health", c.healthController.HealthCheck)
	c.userController.SetupRoutes(router)
	c.mentorRecruitmentController.SetupRoutes(router)
	c.planController.SetupRoutes(router)
}
