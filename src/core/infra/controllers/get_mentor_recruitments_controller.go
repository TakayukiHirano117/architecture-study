package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/mentorrecruitmentapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers/dto"
)

type GetMentorRecruitmentsController struct {
	getMentorRecruitmentsAppService mentorrecruitmentapp.GetMentorRecruitmentsAppService
}

func NewGetMentorRecruitmentsController() *GetMentorRecruitmentsController {
	return &GetMentorRecruitmentsController{
		getMentorRecruitmentsAppService: mentorrecruitmentapp.NewGetMentorRecruitmentsAppService(),
	}
}

func (c *GetMentorRecruitmentsController) Exec(ctx *gin.Context) {
	mentorRecruitments, err := c.getMentorRecruitmentsAppService.Exec(ctx.Request.Context())
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	responses := make([]dto.MentorRecruitmentResponse, 0, len(mentorRecruitments))
	for _, mr := range mentorRecruitments {
		responses = append(responses, dto.ToMentorRecruitmentResponse(mr))
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Mentor recruitments fetched successfully", "data": responses})
}
