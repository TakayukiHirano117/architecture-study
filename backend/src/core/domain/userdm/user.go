// Package userdm はユーザー集約のドメインモデルを定義します.
// ユーザーエンティティ、スキル、経歴、認証情報などの値オブジェクトを含みます.
package userdm

import (
	"time"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
)

// User はユーザー集約のルートエンティティです.
// 名前、メールアドレス、パスワード、スキル、経歴、自己紹介を保持します.
type User struct {
	createdAt        time.Time
	updatedAt        time.Time
	selfIntroduction *SelfIntroduction
	id               shared.UUID
	name             UserName
	password         Password
	email            Email
	skills           []Skill
	careers          []Career
}

// NewUser は新規ユーザーを生成します.
// skills は1件以上必須です. バリデーションに失敗した場合はエラーを返します.
func NewUser(id shared.UUID, name UserName, password Password, email Email, skills []Skill, careers []Career, selfIntroduction *SelfIntroduction) (*User, error) {
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

// NewUserByVal はスキル件数のバリデーションを行わずにユーザーを生成します.
// 既存データの復元など、バリデーションをスキップしたい場合に使用します.
func NewUserByVal(id shared.UUID, name UserName, password Password, email Email, skills []Skill, careers []Career, selfIntroduction *SelfIntroduction) (*User, error) {
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

// CareerParamIfUpdate は経歴の更新時に受け取るパラメータです.
// ID が nil の場合は新規作成、非 nil の場合は既存経歴の更新を表します.
type CareerParamIfUpdate struct {
	ID        *string
	Detail    string
	StartYear uint16
	EndYear   uint16
}

// SkillParamIfUpdate はスキルの更新時に受け取るパラメータです.
// ID が nil の場合は新規作成、非 nil の場合は既存スキルの更新を表します.
type SkillParamIfUpdate struct {
	ID                *string
	Tag               TagParamIfUpdate
	Evaluation        uint8
	YearsOfExperience uint8
}

// TagParamIfUpdate はタグの更新時に受け取るパラメータです.
// ID が nil の場合は新規作成、非 nil の場合は既存タグの更新を表します.
type TagParamIfUpdate struct {
	ID   *string
	Name string
}

// UpdateProfile はユーザーのプロフィール（名前、メール、スキル、経歴、自己紹介）を更新します.
// 各パラメータは値オブジェクトに変換され、バリデーションされます.
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

		var tagID shared.UUID
		if rs.Tag.ID != nil {
			id, err := shared.NewUUIDByVal(*rs.Tag.ID)
			if err != nil {
				return err
			}
			tagID = id
		} else {
			tagID = shared.NewUUID()
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

// ID はユーザーの一意識別子を返します.
func (u *User) ID() shared.UUID {
	return u.id
}

// Name はユーザー名を返します.
func (u *User) Name() UserName {
	return u.name
}

// Password はハッシュ化されたパスワードを返します.
func (u *User) Password() Password {
	return u.password
}

// Skills はユーザーが持つスキル一覧を返します.
func (u *User) Skills() []Skill {
	return u.skills
}

// Careers はユーザーの経歴一覧を返します.
func (u *User) Careers() []Career {
	return u.careers
}

// Email はメールアドレスを返します.
func (u *User) Email() Email {
	return u.email
}

// SelfIntroduction は自己紹介文を返します. 未設定の場合は nil です.
func (u *User) SelfIntroduction() *SelfIntroduction {
	return u.selfIntroduction
}

// CreatedAt は作成日時を返します.
func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAt は更新日時を返します.
func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
