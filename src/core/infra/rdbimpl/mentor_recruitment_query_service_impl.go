package rdbimpl

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/models"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type MentorRecruitmentQueryServiceImpl struct{}

func NewMentorRecruitmentQueryServiceImpl() *MentorRecruitmentQueryServiceImpl {
	return &MentorRecruitmentQueryServiceImpl{}
}

func (s *MentorRecruitmentQueryServiceImpl) GetMentorRecruitments(ctx context.Context) ([]*mentor_recruitmentdm.MentorRecruitment, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, user_id, title, description, category_id, consultation_type, consultation_method, budget_from, budget_to, application_period, status, created_at, updated_at FROM mentor_recruitments ORDER BY created_at DESC
	`
	rows, err := conn.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mentorRecruitments := []*mentor_recruitmentdm.MentorRecruitment{}
	for rows.Next() {
		var r models.MentorRecruitmentModel
		if err := rows.StructScan(&r); err != nil {
			return nil, err
		}
		id, err := mentor_recruitmentdm.NewMentorRecruitmentIDByVal(r.ID)
		if err != nil {
			return nil, err
		}
		userID, err := shared.NewUUIDByVal(r.UserID)
		if err != nil {
			return nil, err
		}
		categoryID, err := categorydm.NewCategoryIDByVal(r.CategoryID)
		if err != nil {
			return nil, err
		}
		consultationType, err := plandm.NewConsultationTypeByVal(r.ConsultationType)
		if err != nil {
			return nil, err
		}
		consultationMethod, err := mentor_recruitmentdm.NewConsultationMethodByVal(r.ConsultationMethod)
		if err != nil {
			return nil, err
		}
		applicationPeriod, err := mentor_recruitmentdm.NewApplicationPeriodByVal(r.ApplicationPeriod)
		if err != nil {
			return nil, err
		}
		status, err := plandm.NewStatusByVal(r.Status)
		if err != nil {
			return nil, err
		}
		budgetFrom := uint32(r.BudgetFrom)
		budgetTo := uint32(r.BudgetTo)
		createdAt := r.CreatedAt
		updatedAt := r.UpdatedAt

		tags := []tagdm.Tag{}
		for _, tag := range r.Tags {
			tagID, err := shared.NewUUIDByVal(tag.ID)
			if err != nil {
				return nil, err
			}
			tagName, err := tagdm.NewTagNameByVal(tag.Name)
			if err != nil {
				return nil, err
			}
			tag, err := tagdm.NewTagByVal(tagID, tagName)
			if err != nil {
				return nil, err
			}
			tags = append(tags, *tag)
		}

		mentorRecruitment, err := mentor_recruitmentdm.NewMentorRecruitmentByVal(id, userID, r.Title, r.Description, categoryID, consultationType, consultationMethod, budgetFrom, budgetTo, applicationPeriod, status, tags, createdAt, updatedAt)
		if err != nil {
			return nil, err
		}

		mentorRecruitments = append(mentorRecruitments, mentorRecruitment)
	}

	return mentorRecruitments, nil
}
