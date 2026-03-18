package rdbimpl

import (
	"context"

	"github.com/lib/pq"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/models"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type PlanQueryServiceImpl struct{}

func NewPlanQueryServiceImpl() *PlanQueryServiceImpl {
	return &PlanQueryServiceImpl{}
}

func (s *PlanQueryServiceImpl) GetPlans(ctx context.Context) ([]*plandm.Plan, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT 
			p.id, 
			p.mentor_id, 
			p.title, 
			p.category_id, 
			p.description, 
			p.status, 
			p.consultation_type, 
			p.price, 
			COALESCE(array_agg(t.id) FILTER (WHERE t.id IS NOT NULL), '{}') AS tag_ids,
			p.created_at, 
			p.updated_at 
		FROM 
			plans p
		LEFT JOIN plan_tags pt ON pt.plan_id = p.id
		LEFT JOIN tags t ON t.id = pt.tag_id
		GROUP BY p.id
		ORDER BY p.created_at DESC
	`
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plans := []*plandm.Plan{}
	for rows.Next() {
		planDetail := models.PlanDetailModel{}
		if err := rows.Scan(
			&planDetail.PlanID,
			&planDetail.MentorID,
			&planDetail.Title,
			&planDetail.CategoryID,
			&planDetail.Description,
			&planDetail.Status,
			&planDetail.ConsultationType,
			&planDetail.Price,
			pq.Array(&planDetail.TagIDs),
			&planDetail.CreatedAt,
			&planDetail.UpdatedAt,
		); err != nil {
			return nil, err
		}

		id, err := shared.NewUUIDByVal(planDetail.PlanID)
		if err != nil {
			return nil, err
		}

		mentorID, err := userdm.NewUserIDByVal(planDetail.MentorID)
		if err != nil {
			return nil, err
		}

		categoryID, err := categorydm.NewCategoryIDByVal(planDetail.CategoryID)
		if err != nil {
			return nil, err
		}

		status, err := plandm.NewStatus(planDetail.Status)
		if err != nil {
			return nil, err
		}

		consultationType, err := plandm.NewConsultationType(planDetail.ConsultationType)
		if err != nil {
			return nil, err
		}

		tagIDs := make([]shared.UUID, 0, len(planDetail.TagIDs))
		for _, tagID := range planDetail.TagIDs {
			uuid, err := shared.NewUUIDByVal(tagID)
			if err != nil {
				return nil, err
			}
			tagIDs = append(tagIDs, uuid)
		}

		plan, err := plandm.NewPlanByVal(
			id,
			mentorID,
			planDetail.Title,
			categoryID,
			tagIDs,
			planDetail.Description,
			status,
			&consultationType,
			uint32(planDetail.Price),
		)
		if err != nil {
			return nil, err
		}

		plans = append(plans, plan)
	}

	return plans, nil
}
