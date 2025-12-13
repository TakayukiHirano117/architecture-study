package models

import "time"

type ContractDetailModel struct {
	ContractID string    `db:"id"`
	MenteeID   string    `db:"user_id"`
	PlanID     string    `db:"plan_id"`
	Message    string    `db:"message"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
