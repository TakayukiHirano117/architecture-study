package dto

import (
	"time"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
)

type MentorRecruitmentResponse struct {
	ID                 string                    `json:"id"`
	UserID             string                    `json:"user_id"`
	Title              string                    `json:"title"`
	Description        string                    `json:"description"`
	CategoryID         string                    `json:"category_id"`
	ConsultationType   string                    `json:"consultation_type"`
	ConsultationMethod string                    `json:"consultation_method"`
	BudgetFrom         uint32                    `json:"budget_from"`
	BudgetTo           uint32                    `json:"budget_to"`
	ApplicationPeriod  time.Time                 `json:"application_period"`
	Status             string                    `json:"status"`
	Tags               []MentorRecruitmentTagDTO `json:"tags"`
	CreatedAt          time.Time                 `json:"created_at"`
	UpdatedAt          time.Time                 `json:"updated_at"`
}

type MentorRecruitmentTagDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToMentorRecruitmentResponse(mr *mentor_recruitmentdm.MentorRecruitment) MentorRecruitmentResponse {
	tags := make([]MentorRecruitmentTagDTO, 0, len(mr.Tags()))
	for _, t := range mr.Tags() {
		tags = append(tags, MentorRecruitmentTagDTO{
			ID:   t.ID().String(),
			Name: t.Name().String(),
		})
	}

	return MentorRecruitmentResponse{
		ID:                 mr.ID().String(),
		UserID:             mr.UserID().String(),
		Title:              mr.Title(),
		Description:        mr.Description(),
		CategoryID:         mr.CategoryID().String(),
		ConsultationType:   string(mr.ConsultationType()),
		ConsultationMethod: string(mr.ConsultationMethod()),
		BudgetFrom:         mr.BudgetFrom(),
		BudgetTo:           mr.BudgetTo(),
		ApplicationPeriod:  mr.ApplicationPeriod().Time(),
		Status:             string(mr.Status()),
		Tags:               tags,
		CreatedAt:          mr.CreatedAt(),
		UpdatedAt:          mr.UpdatedAt(),
	}
}
