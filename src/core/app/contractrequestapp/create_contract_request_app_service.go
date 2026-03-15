package contractrequestapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/contract_requestdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/support/customerr"
)

type CreateContractRequestAppService struct {
	contractRequestRepo contract_requestdm.ContractRequestRepository
	planRepo            plandm.PlanRepository
	isExistByUserID     userdm.IsExistByUserIDDomainService
}

type CreateContractRequestRequest struct {
	MenteeID string `json:"mentee_id"`
	Message  string `json:"message"`
	PlanID   string `json:"plan_id"`
}

func NewCreateContractRequestAppService(
	contractRequestRepo contract_requestdm.ContractRequestRepository,
	planRepo plandm.PlanRepository,
	isExistByUserID userdm.IsExistByUserIDDomainService,
) *CreateContractRequestAppService {
	return &CreateContractRequestAppService{
		contractRequestRepo: contractRequestRepo,
		planRepo:            planRepo,
		isExistByUserID:     isExistByUserID,
	}
}

func (app *CreateContractRequestAppService) Exec(ctx context.Context, req *CreateContractRequestRequest) error {
	menteeID, err := shared.NewUUIDByVal(req.MenteeID)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}

	isExistUser, err := app.isExistByUserID.Exec(ctx, menteeID)
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

	plan, err := app.planRepo.FindByID(ctx, planID)
	if err != nil {
		return customerr.InternalWrapf(err, "failed to find plan: %s", err.Error())
	}
	if plan == nil {
		return customerr.NotFound("plan not found")
	}
	if !plan.IsPublished() {
		return customerr.BadRequest("plan is cancelled")
	}

	priceAtRequest := plan.Price()
	contractRequest, err := contract_requestdm.NewContractRequest(
		shared.NewUUID(),
		req.Message,
		menteeID,
		priceAtRequest,
		planID,
		contract_requestdm.Pending,
	)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}

	return app.contractRequestRepo.Store(ctx, contractRequest)
}
