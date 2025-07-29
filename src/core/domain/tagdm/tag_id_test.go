package tagdm

import (
	"testing"

	"github.com/google/uuid"
)

func TestTagId_NewTagId(t *testing.T) {
	tagId := NewTagId()

	// TagIdが空でないことを確認
	if tagId.String() == "" {
		t.Error("NewTagId() should not return empty string")
	}

	// 生成された値が有効なUUIDかどうか確認
	_, err := uuid.Parse(tagId.String())
	if err != nil {
		t.Errorf("NewTagId() should generate valid UUID, got: %s", tagId.String())
	}
}

func TestTagId_NewTagIdByVal_Success(t *testing.T) {
	validUUID := uuid.New().String()

	tagId, err := NewTagIdByVal(validUUID)
	if err != nil {
		t.Errorf("NewTagIdByVal() with valid UUID should not return error, got: %v", err)
	}

	if tagId.String() != validUUID {
		t.Errorf("NewTagIdByVal() should return correct value, expected: %s, got: %s", validUUID, tagId.String())
	}
}

func TestTagId_NewTagIdByVal_EmptyString(t *testing.T) {
	_, err := NewTagIdByVal("")
	if err == nil {
		t.Error("NewTagIdByVal() with empty string should return error")
	}
}

func TestTagId_NewTagIdByVal_InvalidUUID(t *testing.T) {
	invalidUUID := "invalid-uuid-string"

	_, err := NewTagIdByVal(invalidUUID)
	if err == nil {
		t.Error("NewTagIdByVal() with invalid UUID should return error")
	}
}

func TestTagId_String(t *testing.T) {
	validUUID := uuid.New().String()
	tagId, _ := NewTagIdByVal(validUUID)

	if tagId.String() != validUUID {
		t.Errorf("String() should return correct value, expected: %s, got: %s", validUUID, tagId.String())
	}
}

func TestTagId_Equal(t *testing.T) {
	validUUID := uuid.New().String()
	tagId1, _ := NewTagIdByVal(validUUID)
	tagId2, _ := NewTagIdByVal(validUUID)
	tagId3 := NewTagId()

	// 同じ値のTagIdは等しいはず
	if !tagId1.Equal(tagId2) {
		t.Error("Equal() should return true for same UUID values")
	}

	// 異なる値のTagIdは等しくないはず
	if tagId1.Equal(tagId3) {
		t.Error("Equal() should return false for different UUID values")
	}
}
