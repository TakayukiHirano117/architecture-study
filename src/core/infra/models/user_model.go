package models

import "time"

type UserModel struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
	// このあたりもskill, careerのデータモデルに置き換える
	Skills           []SkillModel `db:"skills"`
	Careers          []CareerModel `db:"careers"`
	SelfIntroduction string      `db:"self_introduction"`
	CreatedAt        time.Time   `db:"created_at"`
	UpdatedAt        time.Time   `db:"updated_at"`
}
