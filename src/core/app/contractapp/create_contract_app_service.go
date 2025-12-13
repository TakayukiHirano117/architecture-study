package contractapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/query_service"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/contractdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/support/customerr"
)

type CreateContractAppService struct {
	contractRepo                                        contractdm.ContractRepository
	isExistByUserIDDomainService                        userdm.IsExistByUserIDDomainService
	isExistByPlanIDDomainService                        plandm.IsExistByPlanIDDomainService
	isExistPlanUserCombinationFromContractsQueryService query_service.IsExistPlanUserCombinationFromContractsQueryService
}

func NewCreateContractAppService(
	contractRepo contractdm.ContractRepository,
	isExistByUserIDDomainService userdm.IsExistByUserIDDomainService,
	isExistByPlanIDDomainService plandm.IsExistByPlanIDDomainService,
	isExistPlanUserCombinationFromContractsQueryService query_service.IsExistPlanUserCombinationFromContractsQueryService,
) *CreateContractAppService {
	return &CreateContractAppService{
		contractRepo:                                        contractRepo,
		isExistByUserIDDomainService:                        isExistByUserIDDomainService,
		isExistByPlanIDDomainService:                        isExistByPlanIDDomainService,
		isExistPlanUserCombinationFromContractsQueryService: isExistPlanUserCombinationFromContractsQueryService,
	}
}

type CreateContractRequest struct {
	MenteeID string `json:"mentee_id"`
	PlanID   string `json:"plan_id"`
	Message  string `json:"message"`
}

func (app *CreateContractAppService) Exec(ctx context.Context, req *CreateContractRequest) error {
	menteeID, err := shared.NewUUIDByVal(req.MenteeID)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}

	isExistUser, err := app.isExistByUserIDDomainService.Exec(ctx, menteeID)
	if err != nil {
		return customerr.InternalWrapf(err, "failed to check if user exists: %s", err.Error())
	}

	if !isExistUser {
		return customerr.NotFound("user not found")
	}

	planID, err := shared.NewUUIDByVal(req.PlanID)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}

	isExistPlan, err := app.isExistByPlanIDDomainService.Exec(ctx, planID)
	if err != nil {
		return customerr.InternalWrapf(err, "failed to check if plan exists: %s", err.Error())
	}

	if !isExistPlan {
		return customerr.NotFound("plan not found")
	}

	// plan_idとuser_idの組み合わせが存在する時は保存できない
	isExistPlanUserCombinationFromContracts, err := app.isExistPlanUserCombinationFromContractsQueryService.Exec(ctx, planID, menteeID)
	if err != nil {
		return customerr.InternalWrapf(err, "failed to check if plan user combination exists: %s", err.Error())
	}

	if isExistPlanUserCombinationFromContracts {
		return customerr.BadRequest("plan user combination already exists")
	}

	contract, err := contractdm.NewContract(shared.NewUUID(), menteeID, planID, req.Message)
	if err != nil {
		return customerr.InternalWrapf(err, "%s", err.Error())
	}

	return app.contractRepo.Store(ctx, contract)
}
