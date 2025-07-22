package userdm

type Career struct {
	id                CareerId
	career_detail     CareerDetail
	career_start_year CareerStartYear
	career_end_year   CareerEndYear
}

func NewCareer(id CareerId, career_detail CareerDetail, career_start_year CareerStartYear, career_end_year CareerEndYear) (*Career, error) {
	return &Career{
		id:                id,
		career_detail:     career_detail,
		career_start_year: career_start_year,
		career_end_year:   career_end_year,
	}, nil
}
