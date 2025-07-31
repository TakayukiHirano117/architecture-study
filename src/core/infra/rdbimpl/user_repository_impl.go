package rdbimpl

import (
	"context"
	"errors"
	"time"

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

// データベース用のDTO
type userDTO struct {
	ID               string    `db:"id"`
	Name             string    `db:"name"`
	Email            string    `db:"email"`
	Password         string    `db:"password"`
	SelfIntroduction string    `db:"self_introduction"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

// User ドメインオブジェクトをDTOに変換
func (dto *userDTO) fromDomain(user *userdm.User) {
	dto.ID = user.Id().String()
	dto.Name = string(user.Name())
	dto.Email = string(user.Email())
	dto.Password = string(user.Password())
	dto.SelfIntroduction = string(user.SelfIntroduction())
	dto.CreatedAt = user.CreatedAt()
	dto.UpdatedAt = user.UpdatedAt()
}

var (
	users []*userdm.User
)

// func init() {
// 	// エラーハンドリングを適切に行う
// 	userName, err := userdm.NewUserName("user1")
// 	if err != nil {
// 		panic(err)
// 	}

// 	email, err := userdm.NewEmail("user1@example.com")
// 	if err != nil {
// 		panic(err)
// 	}

// 	password, err := userdm.NewPassword("password123456") // 12文字以上必要
// 	if err != nil {
// 		panic(err)
// 	}

// 	selfIntroduction, err := userdm.NewSelfIntroduction("self introduction 1")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// ダミーのスキルを作成（最低1つ必要）
// 	dummySkill, err := userdm.NewSkill(
// 		userdm.NewSkillId(),
// 		tagdm.NewTagId(),
// 		5, // 評価（例：5段階評価）
// 		2, // 経験年数（例：2年）
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// GenForTestに値型（デリファレンス）で渡す
// 	u1, err := userdm.GenForTest(
// 		userdm.NewUserId(),
// 		*userName, // ポインタを値にデリファレンス
// 		*email,    // ポインタを値にデリファレンス
// 		*password, // ポインタを値にデリファレンス
// 		[]userdm.Skill{*dummySkill},
// 		[]userdm.Career{},
// 		selfIntroduction, // これはすでにポインタなのでそのまま
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	users = []*userdm.User{u1}
// }

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name userdm.UserName) (*userdm.User, error) {
	for _, user := range users {
		if user.Name() == name {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	// users = append(users, user)
	dto := &userDTO{}
	dto.fromDomain(user)
	query := `
		INSERT INTO users (id, name, email, password, self_introduction, created_at, updated_at)
		VALUES (:id, :name, :email, :password, :self_introduction, NOW(), NOW())
	`
	_, err := r.Connect.NamedExecContext(ctx, query, dto)
	if err != nil {
		return err
	}
	return nil
}
