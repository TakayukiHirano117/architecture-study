package userdm_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestSkillId_NewSkillId(t *testing.T) {
	skillId := userdm.NewSkillID()

	if skillId.String() == "" {
		t.Error("NewSkillId() should not return empty string")
	}

	_, err := uuid.Parse(skillId.String())
	if err != nil {
		t.Errorf("NewSkillId() should generate valid UUID, got: %s", skillId.String())
	}
}

func TestSkillId_NewSkillIdByVal_Success(t *testing.T) {
	validUUID := uuid.New().String()

	skillId, err := userdm.NewSkillIDByVal(validUUID)
	if err != nil {
		t.Errorf("NewSkillIdByVal() with valid UUID should not return error, got: %v", err)
	}

	if skillId.String() != validUUID {
		t.Errorf("NewSkillIdByVal() should return correct value, expected: %s, got: %s", validUUID, skillId.String())
	}
}

func TestSkillId_NewSkillIdByVal_EmptyString(t *testing.T) {
	_, err := userdm.NewSkillIDByVal("")
	if err == nil {
		t.Error("NewSkillIdByVal() with empty string should return error")
	}
}

func TestSkillId_NewSkillIdByVal_InvalidUUID(t *testing.T) {
	invalidUUID := "invalid-uuid-string"

	_, err := userdm.NewSkillIDByVal(invalidUUID)
	if err == nil {
		t.Error("NewSkillIdByVal() with invalid UUID should return error")
	}
}

func TestSkillId_String(t *testing.T) {
	validUUID := uuid.New().String()

	skillId, err := userdm.NewSkillIDByVal(validUUID)
	if err != nil {
		t.Errorf("NewSkillIdByVal() with valid value should not return error, got: %v", err)
	}

	if skillId.String() != validUUID {
		t.Errorf("String() should return correct value, expected: %s, got: %s", validUUID, skillId.String())
	}
}

func TestSkillId_Equal(t *testing.T) {
	validUUID := uuid.New().String()

	skillId1, err := userdm.NewSkillIDByVal(validUUID)
	if err != nil {
		t.Errorf("NewSkillIdByVal() with valid value should not return error, got: %v", err)
	}

	skillId2, err := userdm.NewSkillIDByVal(validUUID)
	if err != nil {
		t.Errorf("NewSkillIdByVal() with valid value should not return error, got: %v", err)
	}

	skillId3 := userdm.NewSkillID()

	if !skillId1.Equal(skillId2) {
		t.Error("Equal() should return true for same UUID values")
	}

	if skillId1.Equal(skillId3) {
		t.Error("Equal() should return false for different UUID values")
	}
}
