package rdbimpl

import (
	"context"
	"errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name userdm.UserName) (*userdm.User, error) {
	conn, err := rdb.ExecFromCtx(ctx)

	if err != nil {
		return nil, err
	}

	query := `
		SELECT id FROM users WHERE name = $1
	`
	rows, err := conn.QueryContext(ctx, query, name.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		return &userdm.User{}, nil
	}

	return nil, nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id userdm.UserID) (*userdm.User, error) {
	conn, err := rdb.ExecFromCtx(ctx)

	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, name, email, password, self_introduction, created_at, updated_at FROM users WHERE id = $1
	`
	// スキル・キャリアもとってくる
	skillQuery := `
		SELECT id, tag_id, evaluation, years_of_experience, created_at, updated_at FROM skills WHERE user_id = $1
	`
	careerQuery := `
		SELECT id, detail, start_year, end_year, created_at, updated_at FROM careers WHERE user_id = $1
	`

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
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return errors.New("transaction not found")
	}

	userQuery := `
		INSERT INTO users (id, name, email, password, self_introduction, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`
	_, err = conn.ExecContext(ctx, userQuery,
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
			_, err = conn.ExecContext(ctx, careerQuery,
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
			_, err = conn.ExecContext(ctx, skillQuery,
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

func (r *UserRepositoryImpl) Update(ctx context.Context, user *userdm.User) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return errors.New("transaction not found")
	}
}
