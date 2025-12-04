//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/categorydm/is_exist_by_category_id_domain_service_mock.go -package=categorydm_mock
package categorydm

import "context"

type IsExistByCategoryIDDomainService interface {
	Exec(ctx context.Context, categoryID CategoryID) (bool, error)
}

type isExistByCategoryIDDomainService struct {
	categoryRepo CategoryRepository
}

func NewIsExistByCategoryIDDomainService(cr CategoryRepository) IsExistByCategoryIDDomainService {
	return &isExistByCategoryIDDomainService{
		categoryRepo: cr,
	}
}

func (iebcidds *isExistByCategoryIDDomainService) Exec(ctx context.Context, categoryID CategoryID) (bool, error) {
	category, err := iebcidds.categoryRepo.FindByID(ctx, categoryID)
	if err != nil {
		return false, err
	}

	return category != nil, nil
}