package userdm

import (
	"regexp"
	"unicode/utf8"
	"strings"

	"github.com/cockroachdb/errors"
)

type Email string

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, errors.New("Email is empty")
	}

	if utf8.RuneCountInString(value) > 255 {
		return nil, errors.New("Email is too long")
	}

	if !strings.Contains(value, "@") {
		return nil, errors.New("Email is invalid")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(value) {
		return nil, errors.New("Email format is invalid")
	}

	email := Email(value)

	return &email, nil
}

func NewEmailByVal(value string) Email {
	return Email(value)
}

func (e Email) String() string {
	return string(e)
}

func (e Email) Equal(e2 Email) bool {
	return e == e2
}
