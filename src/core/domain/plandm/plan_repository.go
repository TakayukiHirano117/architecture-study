//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/plandm/plan_repository_mock.go -package=plandm_mock
package plandm

import "context"

type PlanRepository interface {
	Store(ctx context.Context, plan *Plan) error
}
