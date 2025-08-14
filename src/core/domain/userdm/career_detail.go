package userdm

import (
	"unicode/utf8"

	"github.com/cockroachdb/errors"
)

type CareerDetail string

func NewCareerDetail(value string) (*CareerDetail, error) {
	if value == "" {
		return nil, errors.New("CareerDetail is empty")
	}

	if utf8.RuneCountInString(value) > 2000 {
		return nil, errors.New("CareerDetail is too long")
	}

	careerDetail := CareerDetail(value)

	return &careerDetail, nil
}

func NewCareerDetailByVal(value string) CareerDetail {
	return CareerDetail(value)
}

func (cd CareerDetail) String() string {
	return string(cd)
}

func (cd CareerDetail) Equal(cd2 CareerDetail) bool {
	return cd == cd2
}