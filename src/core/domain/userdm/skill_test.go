package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestSkill_NewSkill_Success(t *testing.T) {
	skillId := userdm.NewSkillID()

	tagId := tagdm.NewTagID()

	tagName, err := tagdm.NewTagName("test-tag-name")
	if err != nil {
		t.Errorf("NewTagName() with valid parameters should not return error, got: %v", err)
	}

	tag, err := tagdm.NewTag(tagId, *tagName)
	if err != nil {
		t.Errorf("NewTag() with valid parameters should not return error, got: %v", err)
	}

	evaluation, err := userdm.NewEvaluation(5)
	if err != nil {
		t.Errorf("NewEvaluation() with valid parameters should not return error, got: %v", err)
	}

	yearsOfExperience, err := userdm.NewYearsOfExperience(3)
	if err != nil {
		t.Errorf("NewYearsOfExperience() with valid parameters should not return error, got: %v", err)
	}

	skill, err := userdm.NewSkill(skillId, tag, evaluation, yearsOfExperience)
	if err != nil {
		t.Errorf("NewSkill() with valid parameters should not return error, got: %v", err)
	}

	if skill == nil {
		t.Error("NewSkill() should not return nil")
	}
}
