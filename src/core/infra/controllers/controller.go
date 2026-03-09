package controllers

import (
	"github.com/gin-gonic/gin"

	contractcontroller "github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers/contract"
	contractrequest_controllers "github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers/contract_request"
)

type Controller struct {
	healthController            *HealthController
	userController              *UserController
	mentorRecruitmentController *MentorRecruitmentController
	planController              *PlanController
	contractController          *contractcontroller.ContractController
	contractRequestController   *contractrequest_controllers.ContractRequestController
}

func NewController() *Controller {
	healthController := NewHealthController()
	userController := NewUserController()
	mentorRecruitmentController := NewMentorRecruitmentController()
	planController := NewPlanController()
	contractController := contractcontroller.NewContractController()
	contractRequestController := contractrequest_controllers.NewContractRequestController()

	return &Controller{
		healthController:            healthController,
		userController:              userController,
		mentorRecruitmentController: mentorRecruitmentController,
		planController:              planController,
		contractController:          contractController,
		contractRequestController:   contractRequestController,
	}
}

func (c *Controller) SetupRoutes(router *gin.Engine) {
	router.GET("/health", c.healthController.HealthCheck)
	c.userController.SetupRoutes(router)
	c.mentorRecruitmentController.SetupRoutes(router)
	c.planController.SetupRoutes(router)
	c.contractController.SetupRoutes(router)
	c.contractRequestController.SetupRoutes(router)
}
