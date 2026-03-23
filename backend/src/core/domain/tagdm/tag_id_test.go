package tagdm

import (
	"testing"

	"github.com/google/uuid"
)

func TestTagId_NewTagId(t *testing.T) {
	tagId := NewTagID()

	if tagId.String() == "" {
		t.Error("NewTagId() should not return empty string")
	}

	_, parseErr := uuid.Parse(tagId.String())
	if parseErr != nil {
		t.Errorf("NewTagId() should generate valid UUID, got: %s", tagId.String())
	}
}

func TestTagId_NewTagIdByVal_Success(t *testing.T) {
	validUUID := uuid.New().String()

	tagId, err := NewTagIDByVal(validUUID)
	if err != nil {
		t.Errorf("NewTagIdByVal() with valid UUID should not return error, got: %v", err)
	}

	if tagId.String() != validUUID {
		t.Errorf("NewTagIdByVal() should return correct value, expected: %s, got: %s", validUUID, tagId.String())
	}
}

func TestTagId_NewTagIdByVal_EmptyString(t *testing.T) {
	_, err := NewTagIDByVal("")
	if err == nil {
		t.Error("NewTagIdByVal() with empty string should return error")
	}
}

func TestTagId_NewTagIdByVal_InvalidUUID(t *testing.T) {
	invalidUUID := "invalid-uuid-string"

	_, err := NewTagIDByVal(invalidUUID)
	if err == nil {
		t.Error("NewTagIdByVal() with invalid UUID should return error")
	}
}

func TestTagId_String(t *testing.T) {
	validUUID := uuid.New().String()

	tagId, err := NewTagIDByVal(validUUID)
	if err != nil {
		t.Errorf("NewTagIdByVal() with valid UUID should not return error, got: %v", err)
	}

	if tagId.String() != validUUID {
		t.Errorf("String() should return correct value, expected: %s, got: %s", validUUID, tagId.String())
	}
}

func TestTagId_Equal(t *testing.T) {
	validUUID1 := uuid.New().String()
	tagId1, err := NewTagIDByVal(validUUID1)
	if err != nil {
		t.Errorf("NewTagIdByVal() with valid UUID should not return error, got: %v", err)
	}

	tagId2, err := NewTagIDByVal(validUUID1)
	if err != nil {
		t.Errorf("NewTagIdByVal() with valid UUID should not return error, got: %v", err)
	}

	tagId3 := NewTagID()

	if !tagId1.Equal(tagId2) {
		t.Error("Equal() should return true for same UUID values")
	}

	if tagId1.Equal(tagId3) {
		t.Error("Equal() should return false for different UUID values")
	}
}
