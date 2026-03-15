//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/contract_requestdm/contact_request_repository_mock.go -package=contract_requestdm_mock
package contract_requestdm

import (
	"context"
)

type ContractRequestRepository interface {
	Store(ctx context.Context, contractRequest *ContractRequest)	error
}
