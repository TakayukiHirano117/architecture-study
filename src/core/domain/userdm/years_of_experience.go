package userdm

import "github.com/cockroachdb/errors"

type YearsOfExperience int

func NewYearsOfExperience(value int) (YearsOfExperience, error) {
	if value < 0 {
		return 0, errors.New("YearsOfExperience is invalid")
	}

	if value > 100 {
		return 0, errors.New("YearsOfExperience is too large")
	}

	return YearsOfExperience(value), nil
}


func NewYearsOfExperienceByVal(value int) (YearsOfExperience, error) {
	return YearsOfExperience(value), nil
}

func (y YearsOfExperience) Int() int {
	return int(y)
}

func (y YearsOfExperience) Equal(y2 YearsOfExperience) bool {
	return y == y2
}