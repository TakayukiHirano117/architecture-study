//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/userdm/is_exist_by_user_name_domain_service_mock.go -package=userdm_mock
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

func (iebunds *isExistByUserNameDomainService) Exec(ctx context.Context, userName UserName) (bool, error) {
	user, err := iebunds.userRepo.FindByName(ctx, userName)

	if err != nil {
		return false, err
	}

	if user == nil {
		return false, nil
	}

	return true, nil
}
