package rdbimpl

import (
	"context"
	"database/sql"
	"time"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepositoryImpl(db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

// Store はユーザーをデータベースに保存します
// func (r *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
// 	query := `
// 		INSERT INTO users (id, name, email, password, self_introduction, created_at, updated_at)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7)
// 	`

// 	_, err := r.db.ExecContext(ctx, query,
// 		user.Id().String(),
// 		string(user.Name()),
// 		string(user.Email()),
// 		string(user.Password()),
// 		string(user.SelfIntroduction()),
// 		user.CreatedAt(),
// 		user.UpdatedAt(),
// 	)

// 	if err != nil {
// 		return errors.Wrap(err, "failed to store user")
// 	}

// 	// スキルの保存処理
// 	if err := r.storeUserSkills(ctx, user); err != nil {
// 		return errors.Wrap(err, "failed to store user skills")
// 	}

// 	// 経歴の保存処理
// 	if err := r.storeUserCareers(ctx, user); err != nil {
// 		return errors.Wrap(err, "failed to store user careers")
// 	}

// 	return nil
// }

// FindByName は名前でユーザーを検索します
func (r *UserRepositoryImpl) FindByName(ctx context.Context, name userdm.UserName) (*userdm.User, error) {
	query := `
		SELECT id, name, email, password, self_introduction, created_at, updated_at
		FROM users
		WHERE name = $1
	`

	var userRow struct {
		ID               string    `db:"id"`
		Name             string    `db:"name"`
		Email            string    `db:"email"`
		Password         string    `db:"password"`
		SelfIntroduction string    `db:"self_introduction"`
		CreatedAt        time.Time `db:"created_at"`
		UpdatedAt        time.Time `db:"updated_at"`
	}

	err := r.db.GetContext(ctx, &userRow, query, string(name))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errors.Wrap(err, "failed to find user by name")
	}

	// ドメインオブジェクトの再構築
	userId, err := userdm.NewUserIdByVal(userRow.ID)
	if err != nil {
		return nil, errors.Wrap(err, "invalid user id")
	}

	userName, err := userdm.NewUserName(userRow.Name)
	if err != nil {
		return nil, errors.Wrap(err, "invalid user name")
	}

	email, err := userdm.NewEmail(userRow.Email)
	if err != nil {
		return nil, errors.Wrap(err, "invalid email")
	}

	password, err := userdm.NewPassword(userRow.Password)
	if err != nil {
		return nil, errors.Wrap(err, "invalid password")
	}

	selfIntroduction, err := userdm.NewSelfIntroduction(userRow.SelfIntroduction)
	if err != nil {
		return nil, errors.Wrap(err, "invalid self introduction")
	}

	// スキルと経歴の取得（簡略化のため空のスライスで初期化）
	skills := []userdm.Skill{}
	careers := []*userdm.Career{}

	user, err := userdm.NewUser(userId, *userName, *password, skills, careers, *email, selfIntroduction)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user domain object")
	}

	return user, nil
}
