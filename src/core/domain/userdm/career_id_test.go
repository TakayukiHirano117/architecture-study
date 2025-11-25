package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/google/uuid"
)

func TestCareerId_NewCareerId(t *testing.T) {
	careerId := userdm.NewCareerID()

	if careerId.String() == "" {
		t.Error("NewCareerId() should not return empty string")
	}

	_, err := uuid.Parse(careerId.String())
	if err != nil {
		t.Errorf("NewCareerId() should generate valid UUID, got: %s", careerId.String())
	}
}

func TestCareerId_NewCareerIdByVal_Success(t *testing.T) {
	validValue := "test-career-id"

	careerId, err := userdm.NewCareerIDByVal(validValue)
	if err != nil {
		t.Errorf("NewCareerIdByVal() with valid value should not return error, got: %v", err)
	}

	if careerId.String() != validValue {
		t.Errorf("NewCareerIdByVal() should return correct value, expected: %s, got: %s", validValue, careerId.String())
	}
}

func TestCareerId_NewCareerIdByVal_EmptyString(t *testing.T) {
	_, err := userdm.NewCareerIDByVal("")
	if err == nil {
		t.Error("NewCareerIdByVal() with empty string should return error")
	}
}

func TestCareerId_String(t *testing.T) {
	validValue := "test-career-id"
	careerId, err := userdm.NewCareerIDByVal(validValue)
	if err != nil {
		t.Errorf("NewCareerIdByVal() with valid value should not return error, got: %v", err)
	}

	if careerId.String() != validValue {
		t.Errorf("String() should return correct value, expected: %s, got: %s", validValue, careerId.String())
	}
}

func TestCareerId_Equal(t *testing.T) {
	validValue := "test-career-id"
	careerId1, err := userdm.NewCareerIDByVal(validValue)
	if err != nil {
		t.Errorf("NewCareerIdByVal() with valid value should not return error, got: %v", err)
	}
	careerId2, err := userdm.NewCareerIDByVal(validValue)
	if err != nil {
		t.Errorf("NewCareerIdByVal() with valid value should not return error, got: %v", err)
	}
	careerId3 := userdm.NewCareerID()

	if !careerId1.Equal(careerId2) {
		t.Error("Equal() should return true for same values")
	}

	if careerId1.Equal(careerId3) {
		t.Error("Equal() should return false for different values")
	}
}
