package tagdm

import "context"

type TagRepository interface {
	FindById(ctx context.Context, id string) (*Tag, error)
}