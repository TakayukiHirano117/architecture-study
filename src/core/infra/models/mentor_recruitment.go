package models

import (
	"time"
)

type MentorRecruitmentModel struct {
	ID                 string     `db:"id"`
	UserID             string     `db:"user_id"`
	Title              string     `db:"title"`
	Description        string     `db:"description"`
	CategoryID         string     `db:"category_id"`
	ConsultationType   string     `db:"consultation_type"`
	ConsultationMethod string     `db:"consultation_method"`
	BudgetFrom         int        `db:"budget_from"`
	BudgetTo           int        `db:"budget_to"`
	ApplicationPeriod  time.Time  `db:"application_period"`
	Status             string     `db:"status"`
	Tags               []TagModel `db:"tags"`
	CreatedAt          time.Time  `db:"created_at"`
	UpdatedAt          time.Time  `db:"updated_at"`
}
