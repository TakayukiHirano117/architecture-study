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

func NewSelfIntroductionByVal(value string) SelfIntroduction {
	return SelfIntroduction(value)
}

func (si SelfIntroduction) String() string {
	return string(si)
}

func (si SelfIntroduction) Equal(si2 SelfIntroduction) bool {
	return si == si2
}
