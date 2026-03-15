package rdbimpl

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/contract_requestdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type ContractRequestRepositoryImpl struct{}

func NewContractRequestRepositoryImpl() *ContractRequestRepositoryImpl {
	return &ContractRequestRepositoryImpl{}
}

func (r *ContractRequestRepositoryImpl) Store(ctx context.Context, contractRequest *contract_requestdm.ContractRequest) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO contract_requests (id, message, user_id, price_at_request, plan_id, is_accepted, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
	`
	_, err = conn.ExecContext(ctx, query,
		contractRequest.ID().String(),
		contractRequest.Message(),
		contractRequest.MenteeID().String(),
		contractRequest.PriceAtRequest(),
		contractRequest.PlanID().String(),
		contractRequest.IsAccepted().String(),
	)
	if err != nil {
		return err
	}

	return nil
}
