package userdm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type SkillID string

func NewSkillID() SkillID {
	return SkillID(uuid.New().String())
}

func NewSkillIDByVal(val string) (SkillID, error) {
	if val == "" {
		return "", errors.New("SkillID is empty")
	}

	if _, err := uuid.Parse(val); err != nil {
		return "", errors.New("SkillID is invalid")
	}

	return SkillID(val), nil
}

func (skillId SkillID) String() string {
	return string(skillId)
}

func (skillId SkillID) Equal(skillId2 SkillID) bool {
	return skillId == skillId2
}
