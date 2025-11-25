package userdm

//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/userdm/is_exist_by_user_name_exclude_self_domain_service_mock.go -package=userdm

import "context"

type IsExistByUserNameExcludeSelfDomainService interface {
	Exec(ctx context.Context, userName UserName, userID UserID) (bool, error)
}

type isExistByUserNameExcludeSelfDomainService struct {
	userRepo UserRepository
}

func NewIsExistByUserNameExcludeSelfDomainService(ur UserRepository) IsExistByUserNameExcludeSelfDomainService {
	return &isExistByUserNameExcludeSelfDomainService{
		userRepo: ur,
	}
}

func (iebunds *isExistByUserNameExcludeSelfDomainService) Exec(ctx context.Context, userName UserName, userID UserID) (bool, error) {
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
