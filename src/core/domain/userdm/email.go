package userdm

import "github.com/cockroachdb/errors"

type Email string

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, errors.New("Email is empty")
	}

	email := Email(value)

	return &email, nil
}