package userdm

import (
	"time"

	"github.com/cockroachdb/errors"
)

type User struct {
	id               UserId
	name             UserName
	password         Password
	skills           []Skill
	careers          []*Career
	email            Email
	selfIntroduction SelfIntroduction
	createdAt        time.Time
	updatedAt        time.Time
}

func NewUser(id UserId, name UserName, password Password, skills []Skill, careers []*Career, email Email, selfIntroduction *SelfIntroduction) (*User, error) {
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
