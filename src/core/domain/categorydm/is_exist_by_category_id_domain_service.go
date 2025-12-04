//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/categorydm/is_exist_by_category_id_domain_service_mock.go -package=categorydm_mock
package categorydm

import "context"

type IsExistByCategoryIDDomainService interface {
	Exec(ctx context.Context, categoryID CategoryID) (bool, error)
}
