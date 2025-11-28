package rdbimpl

import (
	"context"
	"errors"
	"time"

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
	rows, err := conn.QueryContext(ctx, query, id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("tag id not found")
	}

	var tagIdStr string
	var tagNameStr string
	var createdAt, updatedAt time.Time

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
	rows, err := conn.QueryContext(ctx, query, tagName.String())
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

func (r *TagRepositoryImpl) BulkInsert(ctx context.Context, tags []tagdm.Tag) error {
	if len(tags) == 0 {
		return nil
	}

	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO tags (id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`

	for _, tag := range tags {
		_, err = conn.ExecContext(ctx, query, tag.ID().String(), tag.Name().String(), tag.CreatedAt(), tag.UpdatedAt())
		if err != nil {
			return err
		}
	}

	return nil
}
