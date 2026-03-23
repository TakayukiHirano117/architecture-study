package userdm

import (
	"github.com/cockroachdb/errors"
)

type CareerStartYear uint16

func NewCareerStartYear(value uint16) (*CareerStartYear, error) {
	if value < 1970 {
		return nil, errors.New("CareerStartYear is invalid")
	}

	careerStartYear := CareerStartYear(value)

	return &careerStartYear, nil
}

func NewCareerStartYearByVal(value uint16) (CareerStartYear, error) {
	if value == 0 {
		return 0, errors.New("CareerStartYear must not be empty")
	}

	return CareerStartYear(value), nil
}

func (c CareerStartYear) Uint16() uint16 {
	return uint16(c)
}

func (c CareerStartYear) Equal(c2 CareerStartYear) bool {
	return c == c2
}
