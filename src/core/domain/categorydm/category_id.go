package categorydm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type CategoryID string

func NewCategoryID() CategoryID {
	return CategoryID(uuid.New().String())
}

func NewCategoryIDByVal(val string) (CategoryID, error) {
	if val == "" {
		return "", errors.New("CategoryID must not be empty")
	}

	if _, err := uuid.Parse(val); err != nil {
		return "", errors.New("CategoryID must be a valid UUID")
	}

	return CategoryID(val), nil
}

func (categoryId CategoryID) String() string {
	return string(categoryId)
}

func (categoryId CategoryID) Equal(categoryId2 CategoryID) bool {
	return categoryId == categoryId2
}
