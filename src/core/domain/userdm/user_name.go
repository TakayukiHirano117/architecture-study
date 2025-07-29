package userdm

import "github.com/cockroachdb/errors"

type UserName string

func NewUserName(value string) (*UserName, error) {
	if value == "" {
		return nil, errors.New("UserName is empty")
	}

	if len(value) > 255 {
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