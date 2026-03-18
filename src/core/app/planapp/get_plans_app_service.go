//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/app/planapp/get_plans_app_service_mock.go -package=planapp_mock
package planapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
)

// GetPlansAppService はプラン一覧取得のアプリケーションサービスインターフェース
type GetPlansAppService interface {
	Exec(ctx context.Context) ([]*plandm.Plan, error)
}

type getPlansAppServiceImpl struct {
	planQueryService PlanQueryService
}

func NewGetPlansAppService() GetPlansAppService {
	return &getPlansAppServiceImpl{
		planQueryService: rdbimpl.NewPlanQueryServiceImpl(),
	}
}

func (s *getPlansAppServiceImpl) Exec(ctx context.Context) ([]*plandm.Plan, error) {
	plans, err := s.planQueryService.GetPlans(ctx)
	if err != nil {
		return nil, err
	}

	return plans, nil
}
