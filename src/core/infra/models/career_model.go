package models

type CareerModel struct {
	CareerID  string `db:"id"`
	Detail    string `db:"detail"`
	StartYear int    `db:"start_year"`
	EndYear   int    `db:"end_year"`
}
