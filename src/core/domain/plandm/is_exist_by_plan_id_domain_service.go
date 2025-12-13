//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/plandm/is_exist_by_plan_id_domain_service_mock.go -package=plandm_mock
package plandm

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

type IsExistByPlanIDDomainService interface {
	Exec(ctx context.Context, planID shared.UUID) (bool, error)
}

type isExistByPlanIDDomainService struct {
	planRepo PlanRepository
}

func NewIsExistByPlanIDDomainService(pr PlanRepository) IsExistByPlanIDDomainService {
	return &isExistByPlanIDDomainService{
		planRepo: pr,
	}
}

func (iebpidds *isExistByPlanIDDomainService) Exec(ctx context.Context, planID shared.UUID) (bool, error) {
	plan, err := iebpidds.planRepo.FindByID(ctx, planID)
	if err != nil {
		return false, err
	}

	return plan != nil, nil
}
