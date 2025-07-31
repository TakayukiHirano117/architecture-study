package userdm

import (
	"strconv"

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

func (c CareerEndYear) String() string {
	return strconv.Itoa(int(c))
}

func (c CareerEndYear) Equal(c2 CareerEndYear) bool {
	return c == c2
}