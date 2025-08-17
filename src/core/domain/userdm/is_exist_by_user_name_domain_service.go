package userdm

import "context"

type IsExistByUserNameDomainService interface {
	Exec(ctx context.Context, userName UserName) (bool, error)
}

type isExistByUserNameDomainService struct {
	userRepo UserRepository
}

func NewIsExistByUserNameDomainService(ur UserRepository) IsExistByUserNameDomainService {
	return &isExistByUserNameDomainService{
		userRepo: ur,
	}
}

func (iebus *isExistByUserNameDomainService) Exec(ctx context.Context, userName UserName) (bool, error) {
	user, err := iebus.userRepo.FindByName(ctx, userName)

	if err != nil {
		return false, err
	}

	return user != nil, nil
}