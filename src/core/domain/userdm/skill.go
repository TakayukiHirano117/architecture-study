package userdm

import "github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"

type Skill struct {
	id                SkillId
	tagId             tagdm.TagId
	evaluation        int
	yearsOfExperience int
}

func NewSkill(id SkillId, tagId tagdm.TagId, evaluation int, yearsOfExperience int) (*Skill, error) {
	// 必要なバリデーションかける
	return &Skill{
		id:                id,
		tagId:             tagId,
		evaluation:        evaluation,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func NewSkillByVal(id SkillId, tagId tagdm.TagId, evaluation int, yearsOfExperience int) (*Skill, error) {
	return &Skill{
		id:                id,
		tagId:             tagId,
		evaluation:        evaluation,
		yearsOfExperience: yearsOfExperience,
	}, nil
}