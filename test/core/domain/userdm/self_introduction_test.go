package userdm_test

import (
	"strings"
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestSelfIntroduction_NewSelfIntroduction_Success(t *testing.T) {
	validIntroduction := "こんにちは、よろしくお願いします。"

	selfIntroduction, err := userdm.NewSelfIntroduction(validIntroduction)
	if err != nil {
		t.Errorf("NewSelfIntroduction() with valid introduction should not return error, got: %v", err)
	}

	if selfIntroduction == nil {
		t.Error("NewSelfIntroduction() should not return nil")
	}

	if string(*selfIntroduction) != validIntroduction {
		t.Errorf("NewSelfIntroduction() should return correct value, expected: %s, got: %s", validIntroduction, string(*selfIntroduction))
	}
}

func TestSelfIntroduction_NewSelfIntroduction_EmptyString(t *testing.T) {
	_, err := userdm.NewSelfIntroduction("")
	if err == nil {
		t.Error("NewSelfIntroduction() with empty string should return error")
	}
}

func TestSelfIntroduction_NewSelfIntroduction_TooLong(t *testing.T) {
	tooLongIntroduction := strings.Repeat("a", 2001)

	_, err := userdm.NewSelfIntroduction(tooLongIntroduction)
	if err == nil {
		t.Error("NewSelfIntroduction() with too long introduction should return error")
	}
}

func TestSelfIntroduction_NewSelfIntroduction_MaxLength(t *testing.T) {
	maxLengthIntroduction := strings.Repeat("a", 2000)

	selfIntroduction, err := userdm.NewSelfIntroduction(maxLengthIntroduction)
	if err != nil {
		t.Errorf("NewSelfIntroduction() with max length introduction should not return error, got: %v", err)
	}

	if selfIntroduction == nil {
		t.Error("NewSelfIntroduction() should not return nil")
	}
}
