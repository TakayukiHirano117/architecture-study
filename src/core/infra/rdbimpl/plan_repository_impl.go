package rdbimpl

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type PlanRepositoryImpl struct{}

func NewPlanRepositoryImpl() *PlanRepositoryImpl {
	return &PlanRepositoryImpl{}
}

func (r *PlanRepositoryImpl) Store(ctx context.Context, plan *plandm.Plan) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	fmt.Println("plan", plan)

	// plansテーブルへの挿入
	planQuery := `
		INSERT INTO plans (id, mentor_id, title, category_id, description, status, consultation_type, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
	`
	_, err = conn.ExecContext(
		ctx,
		planQuery,
		plan.ID().String(),
		plan.MentorID().String(),
		plan.Title(),
		plan.CategoryID().String(),
		plan.Description(),
		plan.Status(),
		plan.ConsultationType(),
		plan.Price(),
	)
	if err != nil {
		return err
	}

	// plan_tagsテーブルへの挿入
	if len(plan.TagIDs()) > 0 {
		planTagQuery := `
			INSERT INTO plan_tags (id, plan_id, tag_id, created_at, updated_at)
			VALUES ($1, $2, $3, NOW(), NOW())
		`
		for _, tagID := range plan.TagIDs() {
			_, err = conn.ExecContext(
				ctx,
				planTagQuery,
				uuid.New().String(),
				plan.ID().String(),
				tagID.String(),
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
