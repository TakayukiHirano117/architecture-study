package userdm

import (
	"github.com/cockroachdb/errors"
)

type CareerEndYear int

func NewCareerEndYear(value int) (*CareerEndYear, error) {
	if value < 1970 {
		return nil, errors.New("CareerEndYear is invalid")
	}

	careerEndYear := CareerEndYear(value)

	return &careerEndYear, nil
}

func NewCareerEndYearByVal(value int) (CareerEndYear, error) {
	if value == 0 {
		return 0, errors.New("CareerEndYear must not be empty")
	}
	return CareerEndYear(value), nil
}

func (c CareerEndYear) Int() int {
	return int(c) 
}

func (c CareerEndYear) Equal(c2 CareerEndYear) bool {
	return c == c2
}