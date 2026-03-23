package controllers

import (
	"github.com/gin-gonic/gin"

	contractcontroller "github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers/contract"
	contractrequest_controllers "github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers/contract_request"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/middlewares"
)

type Controller struct {
	healthController            *HealthController
	loginController             *LoginController
	userController              *UserController
	mentorRecruitmentController *MentorRecruitmentController
	planController              *PlanController
	contractController          *contractcontroller.ContractController
	contractRequestController   *contractrequest_controllers.ContractRequestController
}

func NewController() *Controller {
	healthController := NewHealthController()
	loginController := NewLoginController()
	userController := NewUserController()
	mentorRecruitmentController := NewMentorRecruitmentController()
	planController := NewPlanController()
	contractController := contractcontroller.NewContractController()
	contractRequestController := contractrequest_controllers.NewContractRequestController()

	return &Controller{
		healthController:            healthController,
		loginController:             loginController,
		userController:              userController,
		mentorRecruitmentController: mentorRecruitmentController,
		planController:              planController,
		contractController:          contractController,
		contractRequestController:   contractRequestController,
	}
}

func (c *Controller) SetupRoutes(router *gin.Engine) {
	// Public routes (認証不要)
	router.GET("/health", c.healthController.HealthCheck)
	router.POST("/login", c.loginController.Exec)

	// ユーザー登録は認証不要
	c.userController.SetupPublicRoutes(router)

	// 一覧画面はログインしていなくても見れる
	c.planController.SetupPublicRoutes(router)

	// Protected routes (認証必須)
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	c.userController.SetupProtectedRoutes(protected)
	c.planController.SetupProtectedRoutes(protected)
	c.mentorRecruitmentController.SetupRoutes(protected)
	c.contractController.SetupRoutes(protected)
	c.contractRequestController.SetupRoutes(protected)
}
