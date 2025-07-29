package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestPassword_NewPassword_Success(t *testing.T) {
	validPassword := "validPassword0"

	password, err := userdm.NewPassword(validPassword)
	if err != nil {
		t.Errorf("NewPassword() with valid password should not return error, got: %v", err)
	}

	if password == nil {
		t.Error("NewPassword() should not return nil")
	}

	if string(*password) != validPassword {
		t.Errorf("NewPassword() should return correct value, expected: %s, got: %s", validPassword, string(*password))
	}
}

func TestPassword_NewPassword_EmptyString(t *testing.T) {
	_, err := userdm.NewPassword("")
	if err == nil {
		t.Error("NewPassword() with empty string should return error")
	}
}

func TestPassword_NewPassword_TooShort(t *testing.T) {
	shortPassword := "short1A"

	_, err := userdm.NewPassword(shortPassword)
	if err == nil {
		t.Error("NewPassword() with too short password should return error")
	}
}

func TestPassword_NewPassword_NoLetter(t *testing.T) {
	noLetterPassword := "123456789012"

	_, err := userdm.NewPassword(noLetterPassword)
	if err == nil {
		t.Error("NewPassword() without letter should return error")
	}
}

func TestPassword_NewPassword_NoNumber(t *testing.T) {
	noNumberPassword := "validPasswordOnly"

	_, err := userdm.NewPassword(noNumberPassword)
	if err == nil {
		t.Error("NewPassword() without number should return error")
	}
}

func TestPassword_NewPassword_MinLength(t *testing.T) {
	minLengthPassword := "validPassw0rd"

	password, err := userdm.NewPassword(minLengthPassword)
	if err != nil {
		t.Errorf("NewPassword() with min length password should not return error, got: %v", err)
	}

	if password == nil {
		t.Error("NewPassword() should not return nil")
	}
}
