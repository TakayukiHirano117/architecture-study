package userdm

import (
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
)

type CareerParamIfCreate struct {
	Detail          string
	CareerStartYear uint16
	CareerEndYear   uint16
}

type SkillParamIfCreate struct {
	ID                *string
	Tag               TagParamIfCreate
	Evaluation        uint8
	YearsOfExperience uint8
}

type TagParamIfCreate struct {
	ID   *string
	Name string
}

func GenIfCreate(
	userName UserName,
	email Email,
	password Password,
	reqCareers []CareerParamIfCreate,
	reqSkills []SkillParamIfCreate,
	selfIntroduction *SelfIntroduction,
) (*User, error) {
	careers := make([]Career, len(reqCareers))
	for i, rc := range reqCareers {
		cd, err := NewCareerDetail(rc.Detail)
		if err != nil {
			return nil, err
		}

		csy, err := NewCareerStartYear(rc.CareerStartYear)
		if err != nil {
			return nil, err
		}

		cey, err := NewCareerEndYear(rc.CareerEndYear)
		if err != nil {
			return nil, err
		}

		c, err := NewCareer(NewCareerID(), *cd, *csy, *cey)

		if err != nil {
			return nil, err
		}
		careers[i] = *c
	}

	skills := make([]Skill, len(reqSkills))
	for i, rs := range reqSkills {
		var tagID shared.UUID
		if rs.Tag.ID != nil {
			id, err := shared.NewUUIDByVal(*rs.Tag.ID)
			if err != nil {
				return nil, err
			}
			tagID = id
		} else {
			tagID = shared.NewUUID()
		}

		tagName, err := tagdm.NewTagNameByVal(rs.Tag.Name)
		if err != nil {
			return nil, err
		}
		tag, err := tagdm.NewTagByVal(tagID, tagName)
		if err != nil {
			return nil, err
		}

		evaluationVo, err := NewEvaluationByVal(rs.Evaluation)
		if err != nil {
			return nil, err
		}

		yearsOfExperienceVo, err := NewYearsOfExperienceByVal(rs.YearsOfExperience)
		if err != nil {
			return nil, err
		}

		s, err := NewSkill(NewSkillID(), tag, evaluationVo, yearsOfExperienceVo)
		if err != nil {
			return nil, err
		}

		skills[i] = *s
	}

	return NewUser(shared.NewUUID(), userName, password, email, skills, careers, selfIntroduction)
}

func GenForTest(id shared.UUID, name UserName, email Email, password Password, skills []Skill, careers []Career, selfIntroduction *SelfIntroduction) (*User, error) {
	return NewUser(id, name, password, email, skills, careers, selfIntroduction)
}
