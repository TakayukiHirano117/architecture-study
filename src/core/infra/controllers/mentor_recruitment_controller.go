package controllers

import (
	"github.com/gin-gonic/gin"
)

type MentorRecruitmentController struct {
	createMentorRecruitmentController *CreateMentorRecruitmentController
	getMentorRecruitmentsController   *GetMentorRecruitmentsController
}

func NewMentorRecruitmentController() *MentorRecruitmentController {
	createMentorRecruitmentController := NewCreateMentorRecruitmentController()
	getMentorRecruitmentsController := NewGetMentorRecruitmentsController()

	return &MentorRecruitmentController{
		createMentorRecruitmentController: createMentorRecruitmentController,
		getMentorRecruitmentsController:   getMentorRecruitmentsController,
	}
}

func (mc *MentorRecruitmentController) SetupRoutes(router *gin.Engine) {
	router.POST("/mentor-recruitments", mc.createMentorRecruitmentController.Exec)
	router.GET("/mentor-recruitments", mc.getMentorRecruitmentsController.Exec)
}
