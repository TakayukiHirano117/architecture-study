package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/planapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
)

type CreatePlanController struct {
	planRepo            plandm.PlanRepository
	buildTags           tagdm.BuildTagsDomainService
	isExistByCategoryID categorydm.IsExistByCategoryIDDomainService
	isExistByUserID     userdm.IsExistByUserIDDomainService
}

func NewCreatePlanController() *CreatePlanController {
	return &CreatePlanController{
		planRepo:            rdbimpl.NewPlanRepositoryImpl(),
		buildTags:           tagdm.NewBuildTagsDomainService(rdbimpl.NewTagRepositoryImpl()),
		isExistByCategoryID: categorydm.NewIsExistByCategoryIDDomainService(rdbimpl.NewCategoryRepositoryImpl()),
		isExistByUserID:     userdm.NewIsExistByUserIDDomainService(rdbimpl.NewUserRepositoryImpl()),
	}
}

func (c *CreatePlanController) Exec(ctx *gin.Context) {
	var in planapp.CreatePlanRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err := planapp.NewCreatePlanAppService(c.isExistByUserID, c.isExistByCategoryID, c.planRepo, c.buildTags).Exec(ctx.Request.Context(), &in); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Plan created successfully"})
}
