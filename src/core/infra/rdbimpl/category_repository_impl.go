package rdbimpl

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/models"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepositoryImpl() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (r *CategoryRepositoryImpl) FindByID(ctx context.Context, id categorydm.CategoryID) (*categorydm.Category, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, name, created_at, updated_at FROM categories WHERE id = $1
	`

	rows, err := conn.QueryxContext(ctx, query, id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categoryDetailRows := []models.CategoryDetailModel{}
	for rows.Next() {
		var r models.CategoryDetailModel
		if err := rows.StructScan(&r); err != nil {
			return nil, err
		}
		categoryDetailRows = append(categoryDetailRows, r)
	}

	if len(categoryDetailRows) == 0 {
		return nil, nil
	}

	c := categoryDetailRows[0]

	categoryID, err := categorydm.NewCategoryIDByVal(c.CategoryID)
	if err != nil {
		return nil, err
	}

	categoryName, err := categorydm.NewCategoryNameByVal(c.CategoryName)
	if err != nil {
		return nil, err
	}

	return categorydm.NewCategoryByVal(categoryID, categoryName, c.CreatedAt, c.UpdatedAt)
}
