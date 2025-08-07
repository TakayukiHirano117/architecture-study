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

func (s *Skill) Id() SkillId {
	return s.id
}

func (s *Skill) TagId() tagdm.TagId {
	return s.tagId
}

func (s *Skill) Evaluation() int {
	return s.evaluation
}

func (s *Skill) YearsOfExperience() int {
	return s.yearsOfExperience
}