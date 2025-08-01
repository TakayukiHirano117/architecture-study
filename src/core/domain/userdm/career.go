package userdm

type Career struct {
	id        CareerId
	detail    CareerDetail
	startYear CareerStartYear
	endYear   CareerEndYear
}

func NewCareer(id CareerId, detail CareerDetail, startYear CareerStartYear, endYear CareerEndYear) (*Career, error) {
	// 必要なバリデーションかける
	return &Career{
		id:        id,
		detail:    detail,
		startYear: startYear,
		endYear:   endYear,
	}, nil
}

func NewCareerByVal(id CareerId, detail CareerDetail, startYear CareerStartYear, endYear CareerEndYear) (*Career, error) {
	return &Career{
		id:        id,
		detail:    detail,
		startYear: startYear,
		endYear:   endYear,
	}, nil
}
