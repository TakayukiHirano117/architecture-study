package userdm

import "context"

type UserDomainService interface {
	IsExistByUserName(ctx context.Context, userName UserName) (bool, error)
}

type userDomainService struct {
	userRepo UserRepository
}

func NewUserDomainService(ur UserRepository) UserDomainService {
	return &userDomainService{
		userRepo: ur,
	}
}

func (uds *userDomainService) IsExistByUserName(ctx context.Context, userName UserName) (bool, error) {
	user, err := uds.userRepo.FindByName(ctx, userName)

	if err != nil {
		return false, err
	}

	return user != nil, nil
}
