package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestCareer_NewCareer_Success(t *testing.T) {
	careerId := userdm.NewCareerId()
	careerDetail, _ := userdm.NewCareerDetail("Web開発に従事")
	careerStartYear, _ := userdm.NewCareerStartYear(2020)
	careerEndYear, _ := userdm.NewCareerEndYear(2022)

	career, err := userdm.NewCareer(careerId, *careerDetail, *careerStartYear, *careerEndYear)
	if err != nil {
		t.Errorf("NewCareer() with valid parameters should not return error, got: %v", err)
	}

	if career == nil {
		t.Error("NewCareer() should not return nil")
	}
}
