package tagdm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type TagID string

// TODO: これvalue受け取らないかも
func NewTagID(value string) (TagID, error) {
	if value == "" {
		return "", errors.New("TagID is empty")
	}

	if _, err := uuid.Parse(value); err != nil {
		return "", errors.New("TagID is invalid")
	}

	return TagID(value), nil
}

func NewTagIDByVal(val string) (TagID, error) {
	if val == "" {
		return "", errors.New("TagID is empty")
	}

	if _, err := uuid.Parse(val); err != nil {
		return "", errors.New("TagID is invalid")
	}

	return TagID(val), nil
}

func (tagID TagID) String() string {
	return string(tagID)
}

func (tagID TagID) Equal(tagID2 TagID) bool {
	return tagID == tagID2
}
