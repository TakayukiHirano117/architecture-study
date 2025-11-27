package category

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
		return "", errors.New("CategoryID is empty")
	}

	return CategoryID(val), nil
}
