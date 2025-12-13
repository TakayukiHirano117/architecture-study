package rdbimpl

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/contractdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/models"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type ContractRepositoryImpl struct{}

func NewContractRepositoryImpl() *ContractRepositoryImpl {
	return &ContractRepositoryImpl{}
}

func (r *ContractRepositoryImpl) Store(ctx context.Context, contract *contractdm.Contract) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO contracts (id, user_id, plan_id, message, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`
	_, err = conn.ExecContext(ctx, query, contract.ID().String(), contract.MenteeID().String(), contract.PlanID().String(), contract.Message())
	if err != nil {
		return err
	}

	return nil
}

func (r *ContractRepositoryImpl) FindByID(ctx context.Context, planID shared.UUID) (*contractdm.Contract, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, user_id, plan_id, message, created_at, updated_at FROM contracts WHERE plan_id = $1
	`
	rows, err := conn.QueryxContext(ctx, query, planID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	contractDetailRows := []models.ContractDetailModel{}
	for rows.Next() {
		var r models.ContractDetailModel
		if err := rows.StructScan(&r); err != nil {
			return nil, err
		}
		contractDetailRows = append(contractDetailRows, r)
	}

	if len(contractDetailRows) == 0 {
		return nil, nil
	}

	contract := contractDetailRows[0]

	contractID, err := shared.NewUUIDByVal(contract.ContractID)
	if err != nil {
		return nil, err
	}

	menteeID, err := shared.NewUUIDByVal(contract.MenteeID)
	if err != nil {
		return nil, err
	}

	return contractdm.NewContractByVal(contractID, menteeID, planID, contract.Message)
}

func (r *ContractRepositoryImpl) FindContractByPlanIDAndUserID(ctx context.Context, planID shared.UUID, userID shared.UUID) (*contractdm.Contract, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, user_id, plan_id, message, created_at, updated_at FROM contracts WHERE plan_id = $1 AND user_id = $2 AND expired_at IS NULL
	`
	rows, err := conn.QueryxContext(ctx, query, planID.String(), userID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var contract models.ContractDetailModel
	if err := rows.StructScan(&contract); err != nil {
		return nil, err
	}

	contractID, err := shared.NewUUIDByVal(contract.ContractID)
	if err != nil {
		return nil, err
	}

	menteeID, err := shared.NewUUIDByVal(contract.MenteeID)
	if err != nil {
		return nil, err
	}

	return contractdm.NewContractByVal(contractID, menteeID, planID, contract.Message)
}
