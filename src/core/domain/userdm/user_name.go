package userdm

import (
	"unicode/utf8"

	"github.com/cockroachdb/errors"
)

type UserName string

func NewUserName(value string) (*UserName, error) {
	if value == "" {
		return nil, errors.New("UserName is empty")
	}

	if utf8.RuneCountInString(value) > 255 {
		return nil, errors.New("UserName is too long")
	}

	userName := UserName(value)

	return &userName, nil
}

func NewUserNameByVal(val string) (UserName, error) {
	if val == "" {
		return "", errors.New("UserName is empty")
	}

	return UserName(val), nil
}

func (u UserName) String() string {
	return string(u)
}

func (u UserName) Equal(u2 UserName) bool {
	return u == u2
}
