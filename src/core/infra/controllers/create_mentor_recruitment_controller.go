package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/mentorrecruitmentapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
)

type CreateMentorRecruitmentController struct {
	mentorRecruitmentRepo mentor_recruitmentdm.MentorRecruitmentRepository
	tagRepo               tagdm.TagRepository
	userRepo              userdm.UserRepository
	isExistByCategoryID   categorydm.IsExistByCategoryIDDomainService
	isExistByUserID       userdm.IsExistByUserIDDomainService
}

func NewCreateMentorRecruitmentController() *CreateMentorRecruitmentController {
	return &CreateMentorRecruitmentController{
		mentorRecruitmentRepo: rdbimpl.NewMentorRecruitmentRepositoryImpl(),
		tagRepo:               rdbimpl.NewTagRepositoryImpl(),
		userRepo:              rdbimpl.NewUserRepositoryImpl(),
		isExistByCategoryID:   categorydm.NewIsExistByCategoryIDDomainService(rdbimpl.NewCategoryRepositoryImpl()),
		isExistByUserID:       userdm.NewIsExistByUserIDDomainService(rdbimpl.NewUserRepositoryImpl()),
	}
}

func (c *CreateMentorRecruitmentController) Exec(ctx *gin.Context) {
	var in mentorrecruitmentapp.CreateMentorRecruitmentRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err := mentorrecruitmentapp.NewCreateMentorRecruitmentAppService(c.isExistByUserID, c.isExistByCategoryID, c.mentorRecruitmentRepo, c.tagRepo).Exec(ctx.Request.Context(), &in); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Mentor recruitment created successfully"})
}
