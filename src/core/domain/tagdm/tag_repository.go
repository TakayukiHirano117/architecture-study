package tagdm

import "context"

type TagRepository interface {
	FindById(ctx context.Context, id TagId) (*Tag, error)
	FindByTagName(ctx context.Context, tagName TagName) (*TagId, error)
}