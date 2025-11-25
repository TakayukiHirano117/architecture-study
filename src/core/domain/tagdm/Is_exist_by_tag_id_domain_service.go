package tagdm

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type IsExistByTagIDDomainService interface {
	Exec(ctx context.Context, reqTagID string) (bool, error)
}

type isExistByTagIDDomainService struct {
}

func NewIsExistByTagIDDomainService() IsExistByTagIDDomainService {
	return &isExistByTagIDDomainService{}
}

func (iebtidds *isExistByTagIDDomainService) Exec(ctx context.Context, reqTagID string) (bool, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return false, err
	}

	query := `
		SELECT EXISTS(SELECT 1 FROM tags WHERE id = $1)
	`
	var isExist bool
	err = conn.QueryRowxContext(ctx, query, reqTagID).Scan(&isExist)
	if err != nil {
		return false, err
	}

	return isExist, nil
}
