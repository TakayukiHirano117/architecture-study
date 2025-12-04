//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/categorydm/category_repository_mock.go -package=categorydm_mock
package categorydm

import "context"

type CategoryRepository interface {
	FindByID(ctx context.Context, id CategoryID) (*Category, error)
}