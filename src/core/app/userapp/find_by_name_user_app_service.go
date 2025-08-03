package userapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type FindByNameUserAppService struct {
	userRepo userdm.UserRepository
}

func NewFindByNameUserAppService(userRepo userdm.UserRepository) *FindByNameUserAppService {
	return &FindByNameUserAppService{
		userRepo: userRepo,
	}
}

type FindByNameUserRequest struct {
	Name string
}

func (app *FindByNameUserAppService) Exec(ctx context.Context, req *FindByNameUserRequest) (*userdm.User, error) {
	user, err := app.userRepo.FindByName(ctx, userdm.UserName(req.Name))
	if err != nil {
		return nil, err
	}
	return user, nil
}
