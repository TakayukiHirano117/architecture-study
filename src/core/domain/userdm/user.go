package userdm

import (
	"time"

	"github.com/cockroachdb/errors"
)

type User struct {
	id               UserID
	name             UserName
	password         Password
	email            Email
	skills           []Skill
	careers          []Career
	selfIntroduction SelfIntroduction
	createdAt        time.Time
	updatedAt        time.Time
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
		selfIntroduction: *selfIntroduction,
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
		selfIntroduction: *selfIntroduction,
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
	}, nil
}

// TODO: ユーザーのドメインルールを表したメソッドを書く
func (u *User) UpdateProfile(name UserName, email Email, skills []Skill, careers []Career, selfIntroduction *SelfIntroduction) (*User, error) {

	return &User{
		id:               u.id,
		name:             name,
		password:         u.password,
		email:            email,
		skills:           skills,
		careers:          careers,
		selfIntroduction: *selfIntroduction,
		createdAt:        u.createdAt,
		updatedAt:        time.Now(),
	}, nil
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

func (u *User) SelfIntroduction() SelfIntroduction {
	return u.selfIntroduction
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
