package tagdm

import "github.com/cockroachdb/errors"

type TagName string

func NewTagName(value string) (*TagName, error) {
	if value == "" {
		return nil, errors.New("tag name is empty")
	}

	if len(value) > 255 {
		return nil, errors.New("tag name is too long")
	}

	tagName := TagName(value)

	return &tagName, nil
}

func NewTagNameByVal(value string) (TagName, error) {
	if value == "" {
		return "", errors.New("tag name is empty")
	}	

	return TagName(value), nil
}

func (t TagName) String() string {
	return string(t)
}

func (t TagName) Equal(t2 TagName) bool {
	return t == t2
}