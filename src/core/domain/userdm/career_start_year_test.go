package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestCareerStartYear_NewCareerStartYear_Success(t *testing.T) {
	validYear := 2020

	careerStartYear, err := userdm.NewCareerStartYear(validYear)
	if err != nil {
		t.Errorf("NewCareerStartYear() with valid year should not return error, got: %v", err)
	}

	if careerStartYear == nil {
		t.Error("NewCareerStartYear() should not return nil")
	}

	if int(*careerStartYear) != validYear {
		t.Errorf("NewCareerStartYear() should return correct value, expected: %d, got: %d", validYear, int(*careerStartYear))
	}
}

func TestCareerStartYear_NewCareerStartYear_Invalid(t *testing.T) {
	invalidYear := 1969

	_, err := userdm.NewCareerStartYear(invalidYear)
	if err == nil {
		t.Error("NewCareerStartYear() with year < 1970 should return error")
	}
}

func TestCareerStartYear_NewCareerStartYear_MinYear(t *testing.T) {
	minYear := 1970

	careerStartYear, err := userdm.NewCareerStartYear(minYear)
	if err != nil {
		t.Errorf("NewCareerStartYear() with min year should not return error, got: %v", err)
	}

	if careerStartYear == nil {
		t.Error("NewCareerStartYear() should not return nil")
	}
}

func TestCareerStartYear_String(t *testing.T) {
	year := 2020
	careerStartYear, _ := userdm.NewCareerStartYear(year)

	expected := "2020"
	if careerStartYear.String() != expected {
		t.Errorf("String() should return correct value, expected: %s, got: %s", expected, careerStartYear.String())
	}
}

func TestCareerStartYear_Equal(t *testing.T) {
	year := 2020
	careerStartYear1, _ := userdm.NewCareerStartYear(year)
	careerStartYear2, _ := userdm.NewCareerStartYear(year)
	careerStartYear3, _ := userdm.NewCareerStartYear(2021)

	if !careerStartYear1.Equal(careerStartYear2) {
		t.Error("Equal() should return true for same year values")
	}

	if careerStartYear1.Equal(careerStartYear3) {
		t.Error("Equal() should return false for different year values")
	}
}
