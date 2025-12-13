package rdbimpl

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/models"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type PlanRepositoryImpl struct{}

func NewPlanRepositoryImpl() *PlanRepositoryImpl {
	return &PlanRepositoryImpl{}
}

// TODO: plan_idとuser_idの組み合わせが存在する時は保存できない様にする, 契約がすでにあったらだわ。
// TODO: 契約が切れたかどうかとか考慮したほうがいい。is_expiredとか。
// TODO: 1ヶ月経過したら契約切れる、とか。
func (r *PlanRepositoryImpl) Store(ctx context.Context, plan *plandm.Plan) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

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

func (r *PlanRepositoryImpl) FindByID(ctx context.Context, planID shared.UUID) (*plandm.Plan, error) {
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
		LEFT JOIN categories c ON c.id = p.category_id
		LEFT JOIN plan_tags pt ON pt.plan_id = p.id
		LEFT JOIN tags t ON t.id = pt.tag_id
		WHERE p.id = $1
		GROUP BY p.id, c.name
	`
	rows, err := conn.QueryContext(ctx, query, planID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("plan id not found")
	}

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

	return plandm.NewPlanByVal(
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
}
