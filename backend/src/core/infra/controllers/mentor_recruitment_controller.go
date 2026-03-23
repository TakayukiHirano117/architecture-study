package controllers

import "github.com/gin-gonic/gin"

type MentorRecruitmentController struct {
	createMentorRecruitmentController *CreateMentorRecruitmentController
}

func NewMentorRecruitmentController() *MentorRecruitmentController {
	createMentorRecruitmentController := NewCreateMentorRecruitmentController()

	return &MentorRecruitmentController{
		createMentorRecruitmentController: createMentorRecruitmentController,
	}
}

func (mc *MentorRecruitmentController) SetupRoutes(router gin.IRouter) {
	router.POST("/mentor-recruitments", mc.createMentorRecruitmentController.Exec)
}
