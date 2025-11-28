package userdm_test

import (
	"strings"
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestUserName_NewUserName_Success(t *testing.T) {
	validName := "Test User"

	userName, err := userdm.NewUserName(validName)
	if err != nil {
		t.Errorf("NewUserName() with valid name should not return error, got: %v", err)
	}

	if userName == nil {
		t.Error("NewUserName() should not return nil")
	}

	if string(*userName) != validName {
		t.Errorf("NewUserName() should return correct value, expected: %s, got: %s", validName, userName.String())
	}
}

func TestUserName_NewUserName_EmptyString(t *testing.T) {
	_, err := userdm.NewUserName("")
	if err == nil {
		t.Error("NewUserName() with empty string should return error")
	}
}

func TestUserName_NewUserName_TooLong(t *testing.T) {
	tooLongName := strings.Repeat("a", 256)

	_, err := userdm.NewUserName(tooLongName)
	if err == nil {
		t.Error("NewUserName() with too long name should return error")
	}
}

func TestUserName_NewUserName_MaxLength(t *testing.T) {
	maxLengthName := strings.Repeat("a", 255)

	userName, err := userdm.NewUserName(maxLengthName)
	if err != nil {
		t.Errorf("NewUserName() with max length name should not return error, got: %v", err)
	}

	if userName == nil {
		t.Error("NewUserName() should not return nil")
	}
}
