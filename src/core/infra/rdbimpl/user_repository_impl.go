package rdbimpl

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/TakayukiHirano117/architecture-study/config"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type UserRepositoryImpl struct {
	Connect *sqlx.DB
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	dbConfig := config.NewDBConfig()
	db, err := dbConfig.Connect()

	if err != nil {
		panic(errors.New("failed to connect to database: " + err.Error()))
	}

	return &UserRepositoryImpl{Connect: db}
}

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name userdm.UserName) (*userdm.User, error) {
	query := `
		SELECT * FROM users WHERE name = :name
	`
	rows, err := r.Connect.NamedQueryContext(ctx, query, map[string]interface{}{"name": name})
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		return nil, errors.New("user name already exists")
	}

	return nil, nil
}

func (r *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	// トランザクション開始
	tx, err := r.Connect.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// ユーザー情報をusersテーブルに挿入
	userQuery := `
		INSERT INTO users (id, name, email, password, self_introduction, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`
	_, err = tx.ExecContext(ctx, userQuery,
		user.Id().String(),
		user.Name().String(),
		user.Email().String(),
		user.Password().String(),
		user.SelfIntroduction().String(),
	)
	if err != nil {
		return err
	}

	// careersテーブルに経歴情報を挿入
	if len(user.Careers()) > 0 {
		careerQuery := `
			INSERT INTO careers (id, user_id, detail, start_year, end_year, created_at, updated_at)
			VALUES (gen_random_uuid(), $1, $2, $3, $4, NOW(), NOW())
		`
		for _, career := range user.Careers() {
			_, err = tx.ExecContext(ctx, careerQuery,
				user.Id().String(),
				career.Detail().String(),
				career.StartYear().Int(),
				career.EndYear().Int(),
			)
			if err != nil {
				return err
			}
		}
	}

	// skillsテーブルにスキル情報を挿入
	if len(user.Skills()) > 0 {
		skillQuery := `
			INSERT INTO skills (id, user_id, tag_id, created_at, updated_at)
			VALUES (gen_random_uuid(), $1, $2, NOW(), NOW())
		`
		for _, skill := range user.Skills() {
			_, err = tx.ExecContext(ctx, skillQuery,
				user.Id().String(),
				skill.TagId().String(),
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
