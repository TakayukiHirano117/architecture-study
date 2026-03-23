package userapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/shared"
	"github.com/TakayukiHirano117/architecture-study/src/support/customerr"
)

type LoginAppService struct {
	userRepo userdm.UserRepository
}

func NewLoginAppService(userRepo userdm.UserRepository) *LoginAppService {
	return &LoginAppService{
		userRepo: userRepo,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (app *LoginAppService) Exec(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	email, err := userdm.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}

	user, err := app.userRepo.FindByEmail(ctx, *email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, customerr.Unauthorized("invalid email or password")
	}

	if !user.Password().CompareWithPlain(req.Password) {
		return nil, customerr.Unauthorized("invalid email or password")
	}

	token, err := shared.GenerateJWT(user.ID().String())
	if err != nil {
		return nil, customerr.InternalWrapf(err, "failed to generate token")
	}

	return &LoginResponse{
		Token: token,
	}, nil
}
