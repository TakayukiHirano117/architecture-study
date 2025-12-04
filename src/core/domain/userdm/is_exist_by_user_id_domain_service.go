//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/userdm/is_exist_by_user_id_domain_service_mock.go -package=userdm_mock
package userdm

import "context"

type IsExistByUserIDDomainService interface {
	Exec(ctx context.Context, userID UserID) (bool, error)
}

type isExistByUserIDDomainService struct {
	userRepo UserRepository
}

func NewIsExistByUserIDDomainService(ur UserRepository) IsExistByUserIDDomainService {
	return &isExistByUserIDDomainService{
		userRepo: ur,
	}
}

func (iebids *isExistByUserIDDomainService) Exec(ctx context.Context, userID UserID) (bool, error) {
	user, err := iebids.userRepo.FindByID(ctx, userID)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}