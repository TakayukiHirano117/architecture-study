package userdm

import (
	"regexp"

	"github.com/cockroachdb/errors"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash password")
	}

	password := Password(string(hashedPassword))

	return &password, nil
}

func NewPasswordByVal(value string) Password {
	return Password(value)
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Equal(p2 Password) bool {
	return p == p2
}