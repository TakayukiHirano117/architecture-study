package tagdm

import "context"

type TagRepository interface {
	FindByID(ctx context.Context, id TagID) (*Tag, error)
	FindIdByTagName(ctx context.Context, tagName TagName) (*TagID, error)
}