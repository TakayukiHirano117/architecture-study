//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/tagdm/tag_repository_mock.go -package=tagdm_mock
package tagdm

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

type TagRepository interface {
	FindByID(ctx context.Context, id shared.UUID) (*Tag, error)
	FindByIDs(ctx context.Context, ids []shared.UUID) ([]Tag, error)
	FindIdByTagName(ctx context.Context, tagName TagName) (*shared.UUID, error)
	BulkInsert(ctx context.Context, tags []Tag) error
}
