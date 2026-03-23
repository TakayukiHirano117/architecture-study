package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/planapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers/dto"
)

type GetPlansController struct {
	getPlansAppService planapp.GetPlansAppService
}

func NewGetPlansController() *GetPlansController {
	return &GetPlansController{
		getPlansAppService: planapp.NewGetPlansAppService(),
	}
}

func (c *GetPlansController) Exec(ctx *gin.Context) {
	plans, err := c.getPlansAppService.Exec(ctx.Request.Context())
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	responses := make([]dto.PlanResponse, 0, len(plans))
	for _, p := range plans {
		responses = append(responses, dto.ToPlanResponse(p))
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Plans fetched successfully", "data": responses})
}
