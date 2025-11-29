package mentor_recruitmentdm

import (
	"unicode/utf8"

	"github.com/cockroachdb/errors"
)

type Description string

func NewDescription(value string) (Description, error) {
	if value == "" {
		return "", errors.New("description must not be empty")
	}

	if utf8.RuneCountInString(value) > 2000 {
		return "", errors.New("description must be less than 2000 characters")
	}

	return Description(value), nil
}

func NewDescriptionByVal(value string) (Description, error) {
	if value == "" {
		return "", errors.New("description must not be empty")
	}

	return Description(value), nil
}

func (d Description) String() string {
	return string(d)
}

func (d Description) Equal(d2 Description) bool {
	return d == d2
}
