package query_service

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/contractdm"
)

type IsExistPlanUserCombinationFromContractsQueryService interface {
	Exec(ctx context.Context, planID shared.UUID, userID shared.UUID) (bool, error)
}

type isExistPlanUserCombinationFromContractsQueryService struct {
	contractRepo contractdm.ContractRepository
}

func NewIsExistPlanUserCombinationFromContractsQueryService(contractRepo contractdm.ContractRepository) IsExistPlanUserCombinationFromContractsQueryService {
	return &isExistPlanUserCombinationFromContractsQueryService{contractRepo: contractRepo}
}

func (iepucfcs *isExistPlanUserCombinationFromContractsQueryService) Exec(ctx context.Context, planID shared.UUID, userID shared.UUID) (bool, error) {
	contract, err := iepucfcs.contractRepo.FindContractByPlanIDAndUserID(ctx, planID, userID)
	if err != nil {
		return false, err
	}

	return contract != nil, nil
}
