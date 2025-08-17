package models

type CareerModel struct {
	Detail    string `db:"detail"`
	StartYear int    `db:"start_year"`
	EndYear   int    `db:"end_year"`
}
