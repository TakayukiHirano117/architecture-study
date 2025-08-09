package tagdm

import "time"

type Tag struct {
	id        TagId
	name      TagName
	createdAt time.Time
	updatedAt time.Time
}

func NewTag(id TagId, name TagName) (*Tag, error) {
	return &Tag{id: id, name: name, createdAt: time.Now(), updatedAt: time.Now()}, nil
}

func NewTagByVal(id TagId, name TagName) (*Tag, error) {
	return &Tag{id: id, name: name, createdAt: time.Now(), updatedAt: time.Now()}, nil
}
