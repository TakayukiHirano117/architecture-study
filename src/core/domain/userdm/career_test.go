package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestCareer_NewCareer_Success(t *testing.T) {
	careerId := userdm.NewCareerID()
	careerDetail, err := userdm.NewCareerDetail("Web開発に従事")
	if err != nil {
		t.Errorf("NewCareerDetail() with valid detail should not return error, got: %v", err)
	}
	careerStartYear, err := userdm.NewCareerStartYear(2020)
	if err != nil {
		t.Errorf("NewCareerStartYear() with valid year should not return error, got: %v", err)
	}
	careerEndYear, err := userdm.NewCareerEndYear(2022)
	if err != nil {
		t.Errorf("NewCareerEndYear() with valid year should not return error, got: %v", err)
	}

	career, err := userdm.NewCareer(careerId, *careerDetail, *careerStartYear, *careerEndYear)
	if err != nil {
		t.Errorf("NewCareer() with valid parameters should not return error, got: %v", err)
	}

	if career == nil {
		t.Error("NewCareer() should not return nil")
	}
}
