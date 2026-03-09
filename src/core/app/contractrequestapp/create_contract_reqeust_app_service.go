package contractrequestapp

import "github.com/TakayukiHirano117/architecture-study/src/core/domain/contractrequestdm"

type CreateContractRequestAppService struct {
	contractRequestRepo contractrequestdm.ContractRequestRepository
}

type CreateContractRequestRequest struct {
	MenteeID       string `json:"mentee_id"`
	Message        string `json:"message"`
	PlanID         string `json:"plan_id"`
	PriceAtRequest uint32 `json:"price_at_request"`
}

// isAcceptedは未確認で作成する
func NewCreateContractRequestAppService(contractRequestRepo contractrequestdm.ContractRequestRepository) *CreateContractRequestAppService {
	return &CreateContractRequestAppService{contractRequestRepo: contractRequestRepo}
}
