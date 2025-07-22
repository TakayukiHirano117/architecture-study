package userdm

import (
	"strings"

	"github.com/cockroachdb/errors"
)

type Password string

func NewPassword(value string) (*Password, error) {
	if value == "" {
		return nil, errors.New("Password is empty")
	}

	if len(value) < 12 {
		return nil, errors.New("Password is too short")
	}

	if !strings.ContainsAny(value, "a-zA-Z") {
		return nil, errors.New("Password must contain at least one letter")
	}

	if !strings.ContainsAny(value, "0-9") {
		return nil, errors.New("Password must contain at least one number")
	}

	password := Password(value)

	return &password, nil
}
