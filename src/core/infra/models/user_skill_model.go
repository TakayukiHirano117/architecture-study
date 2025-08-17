package models

import "time"

type UserSkillModel struct {
	UserID    string    `db:"user_id"`
	SkillID   string    `db:"skill_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
