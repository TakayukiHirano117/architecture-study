package userdm

import (
	"github.com/cockroachdb/errors"
)

type CareerEndYear uint16

func NewCareerEndYear(value uint16) (*CareerEndYear, error) {
	if value < 1970 {
		return nil, errors.New("CareerEndYear is invalid")
	}

	careerEndYear := CareerEndYear(value)

	return &careerEndYear, nil
}

func NewCareerEndYearByVal(value uint16) (CareerEndYear, error) {
	if value == 0 {
		return 0, errors.New("CareerEndYear must not be empty")
	}
	return CareerEndYear(value), nil
}

func (c CareerEndYear) Uint16() uint16 {
	return uint16(c)
}

func (c CareerEndYear) Equal(c2 CareerEndYear) bool {
	return c == c2
}
