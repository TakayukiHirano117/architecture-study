package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestSkill_NewSkill_Success(t *testing.T) {
	skillId := userdm.NewSkillID()
	tagId, _ := tagdm.NewTagID("test-tag-id")
	evaluation := 5
	yearsOfExperience := 3

	skill, err := userdm.NewSkill(skillId, tagId, evaluation, yearsOfExperience)
	if err != nil {
		t.Errorf("NewSkill() with valid parameters should not return error, got: %v", err)
	}

	if skill == nil {
		t.Error("NewSkill() should not return nil")
	}
}
