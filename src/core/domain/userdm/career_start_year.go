package userdm

import (
	"strconv"

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

func (c *CareerStartYear) String() string {
	return strconv.Itoa(int(*c))
}

func (c *CareerStartYear) Equal(c2 *CareerStartYear) bool {
	return *c == *c2
}
