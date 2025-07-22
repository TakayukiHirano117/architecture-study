package userdm

import "github.com/cockroachdb/errors"

type Password string

func NewPassword(value string) (*Password, error) {
	if value == "" {
		return nil, errors.New("Password is empty")
	}

	password := Password(value)

	return &password, nil
}
