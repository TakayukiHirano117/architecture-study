package rdbimpl

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type MentorRecruitmentRepositoryImpl struct{}

func NewMentorRecruitmentRepositoryImpl() *MentorRecruitmentRepositoryImpl {
	return &MentorRecruitmentRepositoryImpl{}
}

func (r *MentorRecruitmentRepositoryImpl) Store(ctx context.Context, mentorRecruitment *mentor_recruitmentdm.MentorRecruitment) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO mentor_recruitments (id, user_id, title, consultation_type, consultation_method, description, budget_from, budget_to, application_period, status, category_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW(), NOW())
	`
	_, err = conn.ExecContext(
		ctx,
		query,
		mentorRecruitment.ID().String(),
		mentorRecruitment.UserID().String(),
		mentorRecruitment.Title(),
		mentorRecruitment.ConsultationType(),
		mentorRecruitment.ConsultationMethod(),
		mentorRecruitment.Description(),
		mentorRecruitment.BudgetFrom(),
		mentorRecruitment.BudgetTo(),
		mentorRecruitment.ApplicationPeriod().Time(),
		mentorRecruitment.Status(),
		mentorRecruitment.CategoryID().String(),
	)
	if err != nil {
		return err
	}

	return nil
}
