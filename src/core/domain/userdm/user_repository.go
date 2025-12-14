//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/userdm/user_repository_mock.go -package=userdm_mock
package userdm

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

type UserRepository interface {
	Store(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindByName(ctx context.Context, name UserName) (*User, error)
	FindByID(ctx context.Context, id shared.UUID) (*User, error)
}
