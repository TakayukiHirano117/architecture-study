package userdm

import "github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"

type CareerParamIfCreate struct {
	Detail          string
	CareerStartYear int
	CareerEndYear   int
}

type SkillParamIfCreate struct {
	TagId             string
	Evaluation        int
	YearsOfExperience int
}

func GenIfCreate(
	userName UserName,
	email Email,
	password Password,
	reqCareers []*CareerParamIfCreate,
	reqSkills []SkillParamIfCreate,
	selfIntroduction *SelfIntroduction,
) (*User, error) {
	careers := make([]*Career, len(reqCareers))
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

		c, err := NewCareer(NewCareerId(), *cd, *csy, *cey)

		if err != nil {
			return nil, err
		}
		careers[i] = c
	}

	skills := make([]Skill, len(reqSkills))
	for i, rs := range reqSkills {
		tagID, err := tagdm.NewTagIdByVal(rs.TagId)

		if err != nil {
			return nil, err
		}

		s, err := NewSkill(NewSkillId(), tagID, rs.Evaluation, rs.YearsOfExperience)

		if err != nil {
			return nil, err
		}

		skills[i] = *s
	}

	return NewUser(NewUserId(), userName, password, email, skills, careers, selfIntroduction)
}
