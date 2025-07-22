package userdm

import "context"

type UserRepository interface {
	Store(ctx context.Context, user *User) error
	FindByName(ctx context.Context, name UserName) (*User, error)
}
