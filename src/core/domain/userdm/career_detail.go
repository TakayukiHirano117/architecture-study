package userdm

import (
	"github.com/cockroachdb/errors"
)

type CareerDetail string

func NewCareerDetail(value string) (*CareerDetail, error) {
	if value == "" {
		return nil, errors.New("CareerDetail is empty")
	}

	if len(value) > 2000 {
		return nil, errors.New("CareerDetail is too long")
	}

	careerDetail := CareerDetail(value)

	return &careerDetail, nil
}
