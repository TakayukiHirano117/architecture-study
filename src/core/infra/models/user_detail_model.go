package models

import "database/sql"

type UserDetailModel struct {
	UserID           string         `db:"user_id"`
	UserName         string         `db:"name"`
	Email            string         `db:"email"`
	Password         string         `db:"password"`
	SelfIntroduction sql.NullString `db:"self_introduction"`

	SkillID                sql.NullString `db:"skill_id"`
	SkillTagID             sql.NullString `db:"skill_tag_id"`
	SkillTagName           sql.NullString `db:"skill_tag_name"`
	SkillEvaluation        sql.NullInt64  `db:"skill_evaluation"`
	SkillYearsOfExperience sql.NullInt64  `db:"skill_years_of_experience"`

	CareerID     sql.NullString `db:"career_id"`
	CareerDetail sql.NullString `db:"career_detail"`
	CareerStart  sql.NullInt64  `db:"career_start_year"`
	CareerEnd    sql.NullInt64  `db:"career_end_year"`
}
