package userdm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type SkillId string

func NewSkillId() SkillId {
	return SkillId(uuid.New().String())
}

func NewSkillIdByVal(val string) (SkillId, error) {
	if val == "" {
		return "", errors.New("SkillId is empty")
	}

	return SkillId(val), nil
}

func (skillId SkillId) String() string {
	return string(skillId)
}

func (skillId SkillId) Equal(skillId2 SkillId) bool {
	return skillId == skillId2
}