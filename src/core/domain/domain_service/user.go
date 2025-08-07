package domain_service

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type UserDomainService interface {
	IsExistByUserName(ctx context.Context, userName userdm.UserName) (bool, error)
}

type userDomainService struct {
	userRepo userdm.UserRepository
}

func NewUserDomainService(ur userdm.UserRepository) UserDomainService {
	return &userDomainService{
		userRepo: ur,
	}
}

func (uds *userDomainService) IsExistByUserName(ctx context.Context, userName userdm.UserName) (bool, error) {
	user, err := uds.userRepo.FindByName(ctx, userName)

	if err != nil {
		return false, err
	}

	return user != nil, nil
}
