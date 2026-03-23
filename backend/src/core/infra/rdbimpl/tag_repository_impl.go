package rdbimpl

import (
	"context"
	"errors"
	"time"

	"github.com/lib/pq"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type TagRepositoryImpl struct {
}

func NewTagRepositoryImpl() *TagRepositoryImpl {
	return &TagRepositoryImpl{}
}

func (r *TagRepositoryImpl) FindByID(ctx context.Context, id shared.UUID) (*tagdm.Tag, error) {
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

	tagID, err := shared.NewUUIDByVal(tagIdStr)
	if err != nil {
		return nil, err
	}

	tagName, err := tagdm.NewTagNameByVal(tagNameStr)
	if err != nil {
		return nil, err
	}

	return tagdm.NewTagByVal(tagID, tagName)
}

func (r *TagRepositoryImpl) FindByIDs(ctx context.Context, ids []shared.UUID) ([]tagdm.Tag, error) {
	if len(ids) == 0 {
		return []tagdm.Tag{}, nil
	}

	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// IDのスライスをstring配列に変換
	idStrs := make([]string, len(ids))
	for i, id := range ids {
		idStrs[i] = id.String()
	}

	query := `
		SELECT id, name FROM tags WHERE id = ANY($1)
	`
	rows, err := conn.QueryContext(ctx, query, pq.Array(idStrs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []tagdm.Tag{}
	for rows.Next() {
		var tagIdStr string
		var tagNameStr string

		if err := rows.Scan(&tagIdStr, &tagNameStr); err != nil {
			return nil, err
		}

		tagID, err := shared.NewUUIDByVal(tagIdStr)
		if err != nil {
			return nil, err
		}

		tagName, err := tagdm.NewTagNameByVal(tagNameStr)
		if err != nil {
			return nil, err
		}

		tag, err := tagdm.NewTagByVal(tagID, tagName)
		if err != nil {
			return nil, err
		}

		tags = append(tags, *tag)
	}

	return tags, nil
}

func (r *TagRepositoryImpl) FindIdByTagName(ctx context.Context, tagName tagdm.TagName) (*shared.UUID, error) {
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

	tagID, err := shared.NewUUIDByVal(idStr)
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
