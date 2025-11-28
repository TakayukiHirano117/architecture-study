package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestYearsOfExperience_NewYearsOfExperience_Success(t *testing.T) {
	var validYearsOfExperience uint8 = 3

	yearsOfExperience, err := userdm.NewYearsOfExperience(validYearsOfExperience)
	if err != nil {
		t.Errorf("NewYearsOfExperience() with valid years of experience should not return error, got: %v", err)
	}

	if yearsOfExperience.Uint8() != validYearsOfExperience {
		t.Errorf("NewYearsOfExperience() should return correct value, expected: %d, got: %d", validYearsOfExperience, yearsOfExperience.Uint8())
	}
}

func TestYearsOfExperience_TooLongYearsOfExperienceReturnError(t *testing.T) {
	var tooLongYearsOfExperience uint8 = 101

	_, err := userdm.NewYearsOfExperience(tooLongYearsOfExperience)
	if err == nil {
		t.Errorf("NewYearsOfExperience() with too long years of experience must be less than 100")
	}
}
