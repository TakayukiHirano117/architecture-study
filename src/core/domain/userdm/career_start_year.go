package userdm

import (
	"github.com/cockroachdb/errors"
)

type CareerStartYear int

func NewCareerStartYear(value int) (*CareerStartYear, error) {
	if value < 1970 {
		return nil, errors.New("CareerStartYear is invalid")
	}

	careerStartYear := CareerStartYear(value)

	return &careerStartYear, nil
}

func NewCareerStartYearByVal(value int) (CareerStartYear, error) {
	if value == 0 {
		return 0, errors.New("CareerStartYear must not be empty")
	}

	return CareerStartYear(value), nil
}

func (c CareerStartYear) Int() int {
	return int(c)
}

func (c CareerStartYear) Equal(c2 CareerStartYear) bool {
	return c == c2
}
