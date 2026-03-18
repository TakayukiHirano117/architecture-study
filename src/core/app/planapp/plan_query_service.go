//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/app/planapp/plan_query_service_mock.go -package=planapp_mock
package planapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
)

type PlanQueryService interface {
	GetPlans(ctx context.Context) ([]*plandm.Plan, error)
}
