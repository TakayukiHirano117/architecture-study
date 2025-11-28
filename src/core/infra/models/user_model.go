package models

import "time"

type UserModel struct {
	CreatedAt        time.Time     `db:"created_at"`
	UpdatedAt        time.Time     `db:"updated_at"`
	ID               string        `db:"id"`
	Name             string        `db:"name"`
	Password         string        `db:"password"`
	Email            string        `db:"email"`
	SelfIntroduction string        `db:"self_introduction"`
	Skills           []SkillModel  `db:"skills"`
	Careers          []CareerModel `db:"careers"`
}
