package userdm_test

import (
	"strings"
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestEmail_NewEmail_Success(t *testing.T) {
	validEmail := "test@example.com"

	email, err := userdm.NewEmail(validEmail)
	if err != nil {
		t.Errorf("NewEmail() with valid email should not return error, got: %v", err)
	}

	if email == nil {
		t.Error("NewEmail() should not return nil")
	}

	if string(*email) != validEmail {
		t.Errorf("NewEmail() should return correct value, expected: %s, got: %s", validEmail, string(*email))
	}
}

func TestEmail_NewEmail_EmptyString(t *testing.T) {
	_, err := userdm.NewEmail("")
	if err == nil {
		t.Error("NewEmail() with empty string should return error")
	}
}

func TestEmail_NewEmail_TooLong(t *testing.T) {
	tooLongEmail := strings.Repeat("a", 250) + "@test.com"

	_, err := userdm.NewEmail(tooLongEmail)
	if err == nil {
		t.Error("NewEmail() with too long email should return error")
	}
}

func TestEmail_NewEmail_NoAtSign(t *testing.T) {
	invalidEmail := "testexample.com"

	_, err := userdm.NewEmail(invalidEmail)
	if err == nil {
		t.Error("NewEmail() without @ should return error")
	}
}

func TestEmail_NewEmail_InvalidFormat(t *testing.T) {
	testCases := []string{
		"test@",
		"@example.com",
		"test@.com",
		"test@com",
	}

	for _, invalidEmail := range testCases {
		_, err := userdm.NewEmail(invalidEmail)
		if err == nil {
			t.Errorf("NewEmail() with invalid format '%s' should return error", invalidEmail)
		}
	}
}
