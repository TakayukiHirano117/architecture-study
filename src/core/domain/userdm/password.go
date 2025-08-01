package userdm

import (
	"regexp"

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

	if matched, _ := regexp.MatchString(`[a-zA-Z]`, value); !matched {
		return nil, errors.New("Password must contain at least one letter")
	}

	if matched, _ := regexp.MatchString(`[0-9]`, value); !matched {
		return nil, errors.New("Password must contain at least one number")
	}

	// ここでハッシュ化する
	password := Password(value)

	return &password, nil
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Equal(p2 Password) bool {
	return p == p2
}