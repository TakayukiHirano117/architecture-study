package models

import "time"

type PlanDetailModel struct {
	PlanID           string    `db:"id"`
	MentorID         string    `db:"mentor_id"`
	Title            string    `db:"title"`
	CategoryID       string    `db:"category_id"`
	Description      string    `db:"description"`
	Status           string    `db:"status"`
	ConsultationType string    `db:"consultation_type"`
	Price            int       `db:"price"`
	TagIDs           []string  `db:"tag_ids"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
