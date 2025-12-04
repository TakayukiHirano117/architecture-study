package userdm

import "github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"

type Skill struct {
	// TODO: tag_idのみを持つ様に改修する
	tag               *tagdm.Tag
	id                SkillID
	evaluation        Evaluation
	yearsOfExperience YearsOfExperience
}

func NewSkill(id SkillID, tag *tagdm.Tag, evaluation Evaluation, yearsOfExperience YearsOfExperience) (*Skill, error) {
	return &Skill{
		id:                id,
		tag:               tag,
		evaluation:        evaluation,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func NewSkillByVal(id SkillID, tag *tagdm.Tag, evaluation Evaluation, yearsOfExperience YearsOfExperience) (*Skill, error) {
	return &Skill{
		id:                id,
		tag:               tag,
		evaluation:        evaluation,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func (s *Skill) ID() SkillID {
	return s.id
}

func (s *Skill) Tag() *tagdm.Tag {
	return s.tag
}

func (s *Skill) TagID() tagdm.TagID {
	return s.tag.ID()
}

func (s *Skill) Evaluation() Evaluation {
	return s.evaluation
}

func (s *Skill) YearsOfExperience() YearsOfExperience {
	return s.yearsOfExperience
}
