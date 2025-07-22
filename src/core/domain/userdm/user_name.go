package userdm

import "github.com/cockroachdb/errors"

type UserName string

func NewUserName(value string) (*UserName, error) {
	if value == "" {
		return nil, errors.New("UserName is empty")
	}

	userName := UserName(value)

	return &userName, nil
}
