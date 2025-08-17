package rdbimpl

import (
	"context"
	"errors"

	"github.com/TakayukiHirano117/architecture-study/config"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/jmoiron/sqlx"
)

type TagRepositoryImpl struct {
	Connect *sqlx.DB
}

func NewTagRepositoryImpl() *TagRepositoryImpl {
	dbConfig := config.NewDBConfig()
	db, err := dbConfig.Connect()

	if err != nil {
		panic(errors.New("failed to connect to database: " + err.Error()))
	}

	return &TagRepositoryImpl{Connect: db}
}

func (r *TagRepositoryImpl) FindByID(ctx context.Context, id tagdm.TagID) (*tagdm.Tag, error) {
	query := `
		SELECT id, name, created_at, updated_at FROM tags WHERE id = $1
	`
	rows, err := r.Connect.QueryContext(ctx, query, id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("tag id not found")
	}

	var tagIdStr string
	var tagNameStr string
	var createdAt, updatedAt interface{} // タイムスタンプは使わないので interface{} で受ける

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
	query := `
		SELECT id FROM tags WHERE name = $1
	`
	rows, err := r.Connect.QueryContext(ctx, query, tagName.String())
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
