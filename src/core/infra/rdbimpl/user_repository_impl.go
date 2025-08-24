package rdbimpl

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/middlewares"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name userdm.UserName) (*userdm.User, error) {
	query := `
		SELECT id FROM users WHERE name = $1
	`
	rows, err := ctx.Value(middlewares.DBKey).(*sqlx.DB).QueryContext(ctx, query, name.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		return &userdm.User{}, nil
	}

	return nil, nil
}

func (r *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	db, ok := ctx.Value(middlewares.DBKey).(*sqlx.Tx)
	if !ok {
		return errors.New("transaction not found")
	}

	userQuery := `
		INSERT INTO users (id, name, email, password, self_introduction, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`
	_, err := db.ExecContext(ctx, userQuery,
		user.ID().String(),
		user.Name().String(),
		user.Email().String(),
		user.Password().String(),
		user.SelfIntroduction().String(),
	)
	if err != nil {
		return err
	}

	if len(user.Careers()) > 0 {
		careerQuery := `
			INSERT INTO careers (id, user_id, detail, start_year, end_year, created_at, updated_at)
			VALUES (gen_random_uuid(), $1, $2, $3, $4, NOW(), NOW())
		`
		for _, career := range user.Careers() {
			_, err := db.ExecContext(ctx, careerQuery,
				user.ID().String(),
				career.Detail().String(),
				career.StartYear().Int(),
				career.EndYear().Int(),
			)
			if err != nil {
				return err
			}
		}
	}

	if len(user.Skills()) > 0 {
		skillQuery := `
			INSERT INTO skills (id, user_id, tag_id, created_at, updated_at)
			VALUES (gen_random_uuid(), $1, $2, NOW(), NOW())
		`
		for _, skill := range user.Skills() {
			_, err := db.ExecContext(ctx, skillQuery,
				user.ID().String(),
				skill.TagID().String(),
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
