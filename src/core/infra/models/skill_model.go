package models

type SkillModel struct {
	SkillID           string `db:"id"`
	TagID             string `db:"tag_id"`
	Evaluation        int    `db:"evaluation"`
	YearsOfExperience int    `db:"years_of_experience"`
}
