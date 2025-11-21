package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestYearsOfExperience_NewYearsOfExperience_Success(t *testing.T) {
	validYearsOfExperience := 3

	yearsOfExperience, err := userdm.NewYearsOfExperience(validYearsOfExperience)
	if err != nil {
		t.Errorf("NewYearsOfExperience() with valid years of experience should not return error, got: %v", err)
	}

	if yearsOfExperience.Int() != validYearsOfExperience {
		t.Errorf("NewYearsOfExperience() should return correct value, expected: %d, got: %d", validYearsOfExperience, yearsOfExperience.Int())
	}
}

func TestYearsOfExperience_NegativeYearsOfExperienceReturnError(t *testing.T) {
	negativeYearsOfExperience := -1

	_, err := userdm.NewYearsOfExperience(negativeYearsOfExperience)
	if err == nil {
		t.Errorf("NewYearsOfExperience() with negative years of experience must be positive")
	}
}

func TestYearsOfExperience_TooLongYearsOfExperienceReturnError(t *testing.T) {
	tooLongYearsOfExperience := 101

	_, err := userdm.NewYearsOfExperience(tooLongYearsOfExperience)
	if err == nil {
		t.Errorf("NewYearsOfExperience() with too long years of experience must be less than 100")
	}
}
