package userdm_test

import (
	"strings"
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestCareerDetail_NewCareerDetail_Success(t *testing.T) {
	validDetail := "Webアプリケーション開発に従事しました。"

	careerDetail, err := userdm.NewCareerDetail(validDetail)
	if err != nil {
		t.Errorf("NewCareerDetail() with valid detail should not return error, got: %v", err)
	}

	if careerDetail == nil {
		t.Error("NewCareerDetail() should not return nil")
	}

	if string(*careerDetail) != validDetail {
		t.Errorf("NewCareerDetail() should return correct value, expected: %s, got: %s", validDetail, string(*careerDetail))
	}
}

func TestCareerDetail_NewCareerDetail_EmptyString(t *testing.T) {
	_, err := userdm.NewCareerDetail("")
	if err == nil {
		t.Error("NewCareerDetail() with empty string should return error")
	}
}

func TestCareerDetail_NewCareerDetail_TooLong(t *testing.T) {
	tooLongDetail := strings.Repeat("a", 2001)

	_, err := userdm.NewCareerDetail(tooLongDetail)
	if err == nil {
		t.Error("NewCareerDetail() with too long detail should return error")
	}
}

func TestCareerDetail_NewCareerDetail_MaxLength(t *testing.T) {
	maxLengthDetail := strings.Repeat("a", 2000)

	careerDetail, err := userdm.NewCareerDetail(maxLengthDetail)
	if err != nil {
		t.Errorf("NewCareerDetail() with max length detail should not return error, got: %v", err)
	}

	if careerDetail == nil {
		t.Error("NewCareerDetail() should not return nil")
	}
}
