//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/tagdm/tag_repository_mock.go -package=tagdm_mock
package tagdm

import "context"

type TagRepository interface {
	FindByID(ctx context.Context, id TagID) (*Tag, error)
	FindByIDs(ctx context.Context, ids []TagID) ([]Tag, error)
	FindIdByTagName(ctx context.Context, tagName TagName) (*TagID, error)
	BulkInsert(ctx context.Context, tags []Tag) error
}
