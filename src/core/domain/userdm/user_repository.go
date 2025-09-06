package userdm

import "context"

type UserRepository interface {
	Store(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindByName(ctx context.Context, name UserName) (*User, error)
	FindByID(ctx context.Context, id UserID) (*User, error)
}
