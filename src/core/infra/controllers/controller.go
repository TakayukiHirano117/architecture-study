package controllers

import (
	"github.com/gin-gonic/gin"

	contractcontroller "github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers/contract"
)

type Controller struct {
	healthController            *HealthController
	userController              *UserController
	mentorRecruitmentController *MentorRecruitmentController
	planController              *PlanController
	contractController          *contractcontroller.ContractController
}

func NewController() *Controller {
	healthController := NewHealthController()
	userController := NewUserController()
	mentorRecruitmentController := NewMentorRecruitmentController()
	planController := NewPlanController()
	contractController := contractcontroller.NewContractController()

	return &Controller{
		healthController:            healthController,
		userController:              userController,
		mentorRecruitmentController: mentorRecruitmentController,
		planController:              planController,
		contractController:          contractController,
	}
}

func (c *Controller) SetupRoutes(router *gin.Engine) {
	router.GET("/health", c.healthController.HealthCheck)
	c.userController.SetupRoutes(router)
	c.mentorRecruitmentController.SetupRoutes(router)
	c.planController.SetupRoutes(router)
	c.contractController.SetupRoutes(router)
}
