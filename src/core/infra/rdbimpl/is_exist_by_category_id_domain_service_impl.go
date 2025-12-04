package rdbimpl

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type isExistByCategoryIDDomainServiceImpl struct {}

func NewIsExistByCategoryIDDomainServiceImpl() *isExistByCategoryIDDomainServiceImpl {
	return &isExistByCategoryIDDomainServiceImpl{}
}

func (iebcidds *isExistByCategoryIDDomainServiceImpl) Exec(ctx context.Context, categoryID categorydm.CategoryID) (bool, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return false, err
	}

	query := `
		SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1)
	`
	var isExist bool
	err = conn.QueryRowxContext(ctx, query, categoryID.String()).Scan(&isExist)
	if err != nil {
		return false, err
	}

	return isExist, nil
}
