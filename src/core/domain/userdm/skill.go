package userdm

import "github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"

type Skill struct {
	id                SkillID
	tagId             tagdm.TagID
	evaluation        Evaluation
	yearsOfExperience YearsOfExperience
}

func NewSkill(id SkillID, tagId tagdm.TagID, evaluation Evaluation, yearsOfExperience YearsOfExperience) (*Skill, error) {
	return &Skill{
		id:                id,
		tagId:             tagId,
		evaluation:        evaluation,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func NewSkillByVal(id SkillID, tagId tagdm.TagID, evaluation Evaluation, yearsOfExperience YearsOfExperience) (*Skill, error) {
	return &Skill{
		id:                id,
		tagId:             tagId,
		evaluation:        evaluation,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func (s *Skill) ID() SkillID {
	return s.id
}

func (s *Skill) TagID() tagdm.TagID {
	return s.tagId
}

func (s *Skill) Evaluation() Evaluation {
	return s.evaluation
}

func (s *Skill) YearsOfExperience() YearsOfExperience {
	return s.yearsOfExperience
}