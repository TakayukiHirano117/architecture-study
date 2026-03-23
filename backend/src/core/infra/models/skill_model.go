package models

type SkillModel struct {
	SkillID           string `db:"id"`
	TagID             string `db:"tag_id"`
	TagName           string `db:"tag_name"`
	Evaluation        uint8  `db:"evaluation"`
	YearsOfExperience uint8  `db:"years_of_experience"`
}
