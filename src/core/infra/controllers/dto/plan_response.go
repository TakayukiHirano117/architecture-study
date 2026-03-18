package dto

import (
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
)

type PlanResponse struct {
	ID               string   `json:"id"`
	MentorID         string   `json:"mentor_id"`
	Title            string   `json:"title"`
	CategoryID       string   `json:"category_id"`
	Description      string   `json:"description"`
	Status           string   `json:"status"`
	ConsultationType string   `json:"consultation_type"`
	Price            uint32   `json:"price"`
	TagIDs           []string `json:"tag_ids"`
}

func ToPlanResponse(p *plandm.Plan) PlanResponse {
	tagIDs := make([]string, 0, len(p.TagIDs()))
	for _, id := range p.TagIDs() {
		tagIDs = append(tagIDs, id.String())
	}

	consultationType := ""
	if p.ConsultationType() != nil {
		consultationType = string(*p.ConsultationType())
	}

	return PlanResponse{
		ID:               p.ID().String(),
		MentorID:         p.MentorID().String(),
		Title:            p.Title(),
		CategoryID:       p.CategoryID().String(),
		Description:      p.Description(),
		Status:           string(p.Status()),
		ConsultationType: consultationType,
		Price:            p.Price(),
		TagIDs:           tagIDs,
	}
}
