//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/plandm/plan_repository_mock.go -package=plandm_mock
package plandm

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

type PlanRepository interface {
	Store(ctx context.Context, plan *Plan) error
	FindByID(ctx context.Context, planID shared.UUID) (*Plan, error)
}
