package userdm

import (
	"strings"

	"github.com/cockroachdb/errors"
)

type Email string

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, errors.New("Email is empty")
	}

	if len(value) > 255 {
		return nil, errors.New("Email is too long")
	}

	if !strings.Contains(value, "@") {
		return nil, errors.New("Email is invalid")
	}

	// 他にも形式のチェックとかをする。

	email := Email(value)

	return &email, nil
}

