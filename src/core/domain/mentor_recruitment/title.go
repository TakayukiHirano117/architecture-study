package mentorrecruitment

import "github.com/cockroachdb/errors"

type Title string

func NewTitle(value string) (Title, error) {
	if value == "" {
		return "", errors.New("title is empty")
	}

	return Title(value), nil
}

func NewTitleByVal(value string) (Title, error) {
	if value == "" {
		return "", errors.New("title must not be empty")
	}

	return Title(value), nil
}

func (t Title) String() string {
	return string(t)
}

func (t Title) Equal(t2 Title) bool {
	return t == t2
}
