package rdbimpl

import (
	"context"
	"errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type TagRepositoryImpl struct {
}

func NewTagRepositoryImpl() *TagRepositoryImpl {
	return &TagRepositoryImpl{}
}

func (r *TagRepositoryImpl) FindByID(ctx context.Context, id tagdm.TagID) (*tagdm.Tag, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, name, created_at, updated_at FROM tags WHERE id = $1
	`
	rows, err := conn.Query(query, id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("tag id not found")
	}

	var tagIdStr string
	var tagNameStr string
	var createdAt, updatedAt interface{}

	err = rows.Scan(&tagIdStr, &tagNameStr, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	tagID, err := tagdm.NewTagIDByVal(tagIdStr)
	if err != nil {
		return nil, err
	}

	tagName, err := tagdm.NewTagNameByVal(tagNameStr)
	if err != nil {
		return nil, err
	}

	return tagdm.NewTagByVal(tagID, tagName)
}

func (r *TagRepositoryImpl) FindIdByTagName(ctx context.Context, tagName tagdm.TagName) (*tagdm.TagID, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id FROM tags WHERE name = $1
	`
	rows, err := conn.Query(query, tagName.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("tag name not found")
	}

	var idStr string
	err = rows.Scan(&idStr)
	if err != nil {
		return nil, err
	}

	tagID, err := tagdm.NewTagIDByVal(idStr)
	if err != nil {
		return nil, err
	}

	return &tagID, nil
}
