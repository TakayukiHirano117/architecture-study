package userdm

import "context"

type IsExistByUserNameDomainService interface {
	Exec(ctx context.Context, userName UserName, userID UserID) (bool, error)
}

type isExistByUserNameDomainService struct {
	userRepo UserRepository
}

func NewIsExistByUserNameDomainService(ur UserRepository) IsExistByUserNameDomainService {
	return &isExistByUserNameDomainService{
		userRepo: ur,
	}
}

func (iebunds *isExistByUserNameDomainService) Exec(ctx context.Context, userName UserName, userID UserID) (bool, error) {
	user, err := iebunds.userRepo.FindByName(ctx, userName)

	if err != nil {
		return false, err
	}

	if user == nil {
		return false, nil
	}

	if user.ID().Equal(userID) {
		return false, nil
	}

	return true, nil
}
