package userdm

import "github.com/cockroachdb/errors"

type SelfIntroduction string

func NewSelfIntroduction(value string) (*SelfIntroduction, error) {
	if value == "" {
		return nil, errors.New("SelfIntroduction is empty")
	}

	if len(value) > 2000 {
		return nil, errors.New("SelfIntroduction is too long")
	}

	selfIntroduction := SelfIntroduction(value)

	return &selfIntroduction, nil
}
