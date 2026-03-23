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

// SetupPublicRoutes は認証不要のルート（一覧）を登録します
func (pc *PlanController) SetupPublicRoutes(router gin.IRouter) {
	router.GET("/plans", pc.getPlansController.Exec)
}

// SetupProtectedRoutes は認証必須のルート（作成）を登録します
func (pc *PlanController) SetupProtectedRoutes(router gin.IRouter) {
	router.POST("/plans", pc.createPlanController.Exec)
}
