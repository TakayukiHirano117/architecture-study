package userdm

import (
	"time"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
)

type User struct {
	createdAt        time.Time
	updatedAt        time.Time
	selfIntroduction *SelfIntroduction
	id               UserID
	name             UserName
	password         Password
	email            Email
	skills           []Skill
	careers          []Career
}

func NewUser(id UserID, name UserName, password Password, email Email, skills []Skill, careers []Career, selfIntroduction *SelfIntroduction) (*User, error) {
	if len(skills) <= 0 {
		return nil, errors.New("skills must be at least 1")
	}

	return &User{
		id:               id,
		name:             name,
		password:         password,
		skills:           skills,
		careers:          careers,
		email:            email,
		selfIntroduction: selfIntroduction,
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
	}, nil
}

func NewUserByVal(id UserID, name UserName, password Password, email Email, skills []Skill, careers []Career, selfIntroduction *SelfIntroduction) (*User, error) {
	return &User{
		id:               id,
		name:             name,
		password:         password,
		skills:           skills,
		careers:          careers,
		email:            email,
		selfIntroduction: selfIntroduction,
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
	}, nil
}

type CareerParamIfUpdate struct {
	ID        *string
	Detail    string
	StartYear uint16
	EndYear   uint16
}
type SkillParamIfUpdate struct {
	ID                *string
	Tag               TagParamIfUpdate
	Evaluation        uint8
	YearsOfExperience uint8
}

type TagParamIfUpdate struct {
	ID   *string
	Name string
}

func (u *User) UpdateProfile(reqUserName string, reqEmail string, reqSkills []SkillParamIfUpdate, reqCareers []CareerParamIfUpdate, reqSelfIntroduction string) error {
	userName, err := NewUserName(reqUserName)
	if err != nil {
		return err
	}

	email, err := NewEmail(reqEmail)
	if err != nil {
		return err
	}

	selfIntroduction, err := NewSelfIntroduction(reqSelfIntroduction)
	if err != nil {
		return err
	}

	careers := make([]Career, len(reqCareers))
	for i, rc := range reqCareers {
		if rc.ID != nil {
			id, err := NewCareerIDByVal(*rc.ID)
			if err != nil {
				return err
			}

			detail, err := NewCareerDetail(rc.Detail)
			if err != nil {
				return err
			}

			startYear, err := NewCareerStartYear(rc.StartYear)
			if err != nil {
				return err
			}

			endYear, err := NewCareerEndYear(rc.EndYear)
			if err != nil {
				return err
			}

			career, err := NewCareer(id, *detail, *startYear, *endYear)
			if err != nil {
				return err
			}

			careers[i] = *career
		} else {
			detail, err := NewCareerDetail(rc.Detail)
			if err != nil {
				return err
			}

			startYear, err := NewCareerStartYear(rc.StartYear)
			if err != nil {
				return err
			}

			endYear, err := NewCareerEndYear(rc.EndYear)
			if err != nil {
				return err
			}

			career, err := NewCareer(NewCareerID(), *detail, *startYear, *endYear)
			if err != nil {
				return err
			}

			careers[i] = *career
		}
	}

	skills := make([]Skill, len(reqSkills))
	for i, rs := range reqSkills {
		// SkillIDの処理：IDがある場合は既存のIDを使用、ない場合は新規生成
		var skillID SkillID
		if rs.ID != nil {
			id, err := NewSkillIDByVal(*rs.ID)
			if err != nil {
				return err
			}
			skillID = id
		} else {
			skillID = NewSkillID()
		}

		var tagID tagdm.TagID
		if rs.Tag.ID != nil {
			id, err := tagdm.NewTagIDByVal(*rs.Tag.ID)
			if err != nil {
				return err
			}
			tagID = id
		} else {
			tagID = tagdm.NewTagID()
		}

		tagName, err := tagdm.NewTagNameByVal(rs.Tag.Name)
		if err != nil {
			return err
		}

		tag, err := tagdm.NewTagByVal(tagID, tagName)
		if err != nil {
			return err
		}

		evaluationVo, err := NewEvaluationByVal(rs.Evaluation)
		if err != nil {
			return err
		}

		yearsOfExperienceVo, err := NewYearsOfExperienceByVal(rs.YearsOfExperience)
		if err != nil {
			return err
		}

		skill, err := NewSkill(skillID, tag, evaluationVo, yearsOfExperienceVo)
		if err != nil {
			return err
		}
		skills[i] = *skill
	}

	u.name = *userName
	u.email = *email
	u.selfIntroduction = selfIntroduction
	u.careers = careers
	u.skills = skills

	return nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() UserName {
	return u.name
}

func (u *User) Password() Password {
	return u.password
}

func (u *User) Skills() []Skill {
	return u.skills
}

func (u *User) Careers() []Career {
	return u.careers
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) SelfIntroduction() *SelfIntroduction {
	return u.selfIntroduction
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
