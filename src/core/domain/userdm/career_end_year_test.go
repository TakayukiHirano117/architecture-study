package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestCareerEndYear_NewCareerEndYear_Success(t *testing.T) {
	validYear := 2020

	careerEndYear, err := userdm.NewCareerEndYear(validYear)
	if err != nil {
		t.Errorf("NewCareerEndYear() with valid year should not return error, got: %v", err)
	}

	if careerEndYear == nil {
		t.Error("NewCareerEndYear() should not return nil")
	}

	if int(*careerEndYear) != validYear {
		t.Errorf("NewCareerEndYear() should return correct value, expected: %d, got: %d", validYear, int(*careerEndYear))
	}
}

func TestCareerEndYear_NewCareerEndYear_Invalid(t *testing.T) {
	invalidYear := 1969

	_, err := userdm.NewCareerEndYear(invalidYear)
	if err == nil {
		t.Error("NewCareerEndYear() with year < 1970 should return error")
	}
}

func TestCareerEndYear_NewCareerEndYear_MinYear(t *testing.T) {
	minYear := 1970

	careerEndYear, err := userdm.NewCareerEndYear(minYear)
	if err != nil {
		t.Errorf("NewCareerEndYear() with min year should not return error, got: %v", err)
	}

	if careerEndYear == nil {
		t.Error("NewCareerEndYear() should not return nil")
	}
}

func TestCareerEndYear_String(t *testing.T) {
	year := 2020
	careerEndYear, err := userdm.NewCareerEndYear(year)
	if err != nil {
		t.Errorf("NewCareerEndYear() with valid year should not return error, got: %v", err)
	}

	if int(*careerEndYear) != year {
		t.Errorf("String() should return correct value, expected: %d, got: %d", year, int(*careerEndYear))
	}
}

func TestCareerEndYear_Equal(t *testing.T) {
	year := 2020
	careerEndYear1, err := userdm.NewCareerEndYear(year)
	if err != nil {
		t.Errorf("NewCareerEndYear() with valid year should not return error, got: %v", err)
	}
	careerEndYear2, err := userdm.NewCareerEndYear(year)
	if err != nil {
		t.Errorf("NewCareerEndYear() with valid year should not return error, got: %v", err)
	}
	careerEndYear3, err := userdm.NewCareerEndYear(2021)
	if err != nil {
		t.Errorf("NewCareerEndYear() with valid year should not return error, got: %v", err)
	}

	if !careerEndYear1.Equal(*careerEndYear2) {
		t.Error("Equal() should return true for same year values")
	}

	if careerEndYear1.Equal(*careerEndYear3) {
		t.Error("Equal() should return false for different year values")
	}
}
