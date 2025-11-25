package tagdm

import "time"

type Tag struct {
	id        TagID
	name      TagName
	createdAt time.Time
	updatedAt time.Time
}

func NewTag(id TagID, name TagName) (*Tag, error) {
	return &Tag{id: id, name: name, createdAt: time.Now(), updatedAt: time.Now()}, nil
}

func NewTagByVal(id TagID, name TagName) (*Tag, error) {
	return &Tag{id: id, name: name, createdAt: time.Now(), updatedAt: time.Now()}, nil
}

func (t *Tag) ID() TagID {
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
