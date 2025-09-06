package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/google/uuid"
)

func TestUserId_NewUserId(t *testing.T) {
	userId := userdm.NewUserID()

	if userId.String() == "" {
		t.Error("NewUserId() should not return empty string")
	}

	_, err := uuid.Parse(userId.String())
	if err != nil {
		t.Errorf("NewUserId() should generate valid UUID, got: %s", userId.String())
	}
}

func TestUserId_NewUserIdByVal_Success(t *testing.T) {
	validUUID := uuid.New().String()

	userId, err := userdm.NewUserIDByVal(validUUID)
	if err != nil {
		t.Errorf("NewUserIdByVal() with valid UUID should not return error, got: %v", err)
	}

	if userId.String() != validUUID {
		t.Errorf("NewUserIdByVal() should return correct value, expected: %s, got: %s", validUUID, userId.String())
	}
}

func TestUserId_NewUserIdByVal_EmptyString(t *testing.T) {
	_, err := userdm.NewUserIDByVal("")
	if err == nil {
		t.Error("NewUserIdByVal() with empty string should return error")
	}
}

func TestUserId_NewUserIdByVal_InvalidUUID(t *testing.T) {
	invalidUUID := "invalid-uuid-string"

	_, err := userdm.NewUserIDByVal(invalidUUID)
	if err == nil {
		t.Error("NewUserIdByVal() with invalid UUID should return error")
	}
}

func TestUserId_String(t *testing.T) {
	validUUID := uuid.New().String()
	userId, _ := userdm.NewUserIDByVal(validUUID)

	if userId.String() != validUUID {
		t.Errorf("String() should return correct value, expected: %s, got: %s", validUUID, userId.String())
	}
}

func TestUserId_Equal(t *testing.T) {
	validUUID := uuid.New().String()
	userId1, _ := userdm.NewUserIDByVal(validUUID)
	userId2, _ := userdm.NewUserIDByVal(validUUID)
	userId3 := userdm.NewUserID()

	if !userId1.Equal(userId2) {
		t.Error("Equal() should return true for same UUID values")
	}

	if userId1.Equal(userId3) {
		t.Error("Equal() should return false for different UUID values")
	}
}
