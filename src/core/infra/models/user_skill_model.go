package models

import "time"

type UserSkillModel struct {
	UserId    string    `db:"user_id"`
	SkillId   string    `db:"skill_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
