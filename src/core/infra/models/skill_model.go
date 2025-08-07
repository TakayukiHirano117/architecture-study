package models

type SkillModel struct {
	TagId             string `db:"tag_id"`
	Evaluation        int    `db:"evaluation"`
	YearsOfExperience int    `db:"years_of_experience"`
}