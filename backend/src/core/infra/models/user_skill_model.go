package models

import "time"

type UserSkillModel struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	UserID    string    `db:"user_id"`
	SkillID   string    `db:"skill_id"`
}
