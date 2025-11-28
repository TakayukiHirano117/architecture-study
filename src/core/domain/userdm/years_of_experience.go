package userdm

import "github.com/cockroachdb/errors"

type YearsOfExperience uint8

func NewYearsOfExperience(value uint8) (YearsOfExperience, error) {
	if value > 100 {
		return 0, errors.New("YearsOfExperience is too large")
	}

	return YearsOfExperience(value), nil
}

func NewYearsOfExperienceByVal(value uint8) (YearsOfExperience, error) {
	return YearsOfExperience(value), nil
}

func (y YearsOfExperience) Uint8() uint8 {
	return uint8(y)
}

func (y YearsOfExperience) Equal(y2 YearsOfExperience) bool {
	return y == y2
}
