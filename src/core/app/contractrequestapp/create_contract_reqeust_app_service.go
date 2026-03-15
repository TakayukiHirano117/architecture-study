package contractrequestapp

import "github.com/TakayukiHirano117/architecture-study/src/core/domain/contract_requestdm"

type CreateContractRequestAppService struct {
	contractRequestRepo contract_requestdm.ContractRequestRepository
}

type CreateContractRequestRequest struct {
	MenteeID       string `json:"mentee_id"`
	Message        string `json:"message"`
	PlanID         string `json:"plan_id"`
	PriceAtRequest uint32 `json:"price_at_request"`
}

// isAcceptedは未確認で作成する
func NewCreateContractRequestAppService(contractRequestRepo contract_requestdm.ContractRequestRepository) *CreateContractRequestAppService {
	return &CreateContractRequestAppService{contractRequestRepo: contractRequestRepo}
}

func (app *CreateContractRequestAppService) Exec(ctx context.Context, req *CreateContractRequestRequest) error {
	menteeID, err := shared.NewUUIDByVal(req.MenteeID)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}
}
