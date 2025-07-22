package tagdm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type TagId string

func NewTagId() TagId {
	return TagId(uuid.New().String())
}

func NewTagIdByVal(val string) (TagId, error) {
	if val == "" {
		return "", errors.New("TagId is empty")
	}

	if _, err := uuid.Parse(val); err != nil {
		return "", errors.New("TagId is invalid")
	}

	return TagId(val), nil
}

func (tagId TagId) String() string {
	return string(tagId)
}

func (tagId TagId) Equal(tagId2 TagId) bool {
	return tagId == tagId2
}