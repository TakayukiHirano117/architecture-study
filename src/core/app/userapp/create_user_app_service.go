package userapp

import (
	// "context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type CreateUserAppService struct {
	userRepo userdm.UserRepository
}

func NewCreateUserAppService(userRepo userdm.UserRepository) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo: userRepo,
	}
}

type CreateUserRequest struct {
	name             string
	email            string
	password         string
	skills           []string
	careers          []string
	selfIntroduction string
}

// func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	
// }
