//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/contractdm/contract_repository_mock.go -package=contractdm_mock
package contractdm

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

type ContractRepository interface {
	Store(ctx context.Context, contract *Contract) error
	FindByID(ctx context.Context, planID shared.UUID) (*Contract, error)
	FindContractByPlanIDAndUserID(ctx context.Context, planID shared.UUID, userID shared.UUID) (*Contract, error)
}
