package models

type CareerModel struct {
	CareerID  string `db:"id"`
	Detail    string `db:"detail"`
	StartYear uint16 `db:"start_year"`
	EndYear   uint16 `db:"end_year"`
}
