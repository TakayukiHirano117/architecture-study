package categorydm

import (
	"unicode/utf8"

	"github.com/cockroachdb/errors"
)

type CategoryName string

func NewCategoryName(value string) (*CategoryName, error) {
	if value == "" {
		return nil, errors.New("CategoryName is empty")
	}

	if utf8.RuneCountInString(value) > 255 {
		return nil, errors.New("CategoryName is too long")
	}

	categoryName := CategoryName(value)

	return &categoryName, nil
}

func NewCategoryNameByVal(value string) (CategoryName, error) {
	if value == "" {
		return "", errors.New("CategoryName must not be empty")
	}

	return CategoryName(value), nil
}

func (categoryName CategoryName) String() string {
	return string(categoryName)
}

func (categoryName CategoryName) Equal(categoryName2 CategoryName) bool {
	return categoryName == categoryName2
}
