package models

type SkillModel struct {
	TagID             string `db:"tag_id"`
	Evaluation        int    `db:"evaluation"`
	YearsOfExperience int    `db:"years_of_experience"`
}