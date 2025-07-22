package userdm

import "github.com/cockroachdb/errors"

type SelfIntroduction string

func NewSelfIntroduction(value string) (*SelfIntroduction, error) {
	if value == "" {
		return nil, errors.New("SelfIntroduction is empty")
	}

	selfIntroduction := SelfIntroduction(value)

	return &selfIntroduction, nil
}
