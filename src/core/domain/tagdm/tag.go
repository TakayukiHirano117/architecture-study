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
