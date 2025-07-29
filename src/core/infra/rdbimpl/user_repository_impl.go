package rdbimpl

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/cockroachdb/errors"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

var (
	users []*userdm.User
)

func init() {
	// エラーハンドリングを適切に行う
	userName, err := userdm.NewUserName("user1")
	if err != nil {
		panic(err)
	}

	email, err := userdm.NewEmail("user1@example.com")
	if err != nil {
		panic(err)
	}

	password, err := userdm.NewPassword("password123456") // 12文字以上必要
	if err != nil {
		panic(err)
	}

	selfIntroduction, err := userdm.NewSelfIntroduction("self introduction 1")
	if err != nil {
		panic(err)
	}

	// GenForTestに値型（デリファレンス）で渡す
	u1, err := userdm.GenForTest(
		userdm.NewUserId(),
		*userName, // ポインタを値にデリファレンス
		*email,    // ポインタを値にデリファレンス
		*password, // ポインタを値にデリファレンス
		[]userdm.Skill{},
		[]*userdm.Career{},
		selfIntroduction, // これはすでにポインタなのでそのまま
	)
	if err != nil {
		panic(err)
	}

	users = []*userdm.User{u1}
}

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name userdm.UserName) (*userdm.User, error) {
	for _, user := range users {
		if user.Name() == name {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	users = append(users, user)
	return nil
}
