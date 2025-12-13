package tagdm

import (
	"time"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

type Tag struct {
	createdAt time.Time
	updatedAt time.Time
	id        shared.UUID
	name      TagName
}

func NewTag(id shared.UUID, name TagName) (*Tag, error) {
	return &Tag{id: id, name: name, createdAt: time.Now(), updatedAt: time.Now()}, nil
}

func NewTagByVal(id shared.UUID, name TagName) (*Tag, error) {
	return &Tag{id: id, name: name, createdAt: time.Now(), updatedAt: time.Now()}, nil
}

func (t *Tag) ID() shared.UUID {
	return t.id
}

func (t *Tag) Name() TagName {
	return t.name
}

func (t *Tag) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Tag) UpdatedAt() time.Time {
	return t.updatedAt
}
