package userdm

type Career struct {
	id        CareerID
	detail    CareerDetail
	startYear CareerStartYear
	endYear   CareerEndYear
}

func NewCareer(id CareerID, detail CareerDetail, startYear CareerStartYear, endYear CareerEndYear) (*Career, error) {
	// 必要なバリデーションかける
	return &Career{
		id:        id,
		detail:    detail,
		startYear: startYear,
		endYear:   endYear,
	}, nil
}

func NewCareerByVal(id CareerID, detail CareerDetail, startYear CareerStartYear, endYear CareerEndYear) (*Career, error) {
	return &Career{
		id:        id,
		detail:    detail,
		startYear: startYear,
		endYear:   endYear,
	}, nil
}

func (c *Career) ID() CareerID {
	return c.id
}

func (c *Career) Detail() CareerDetail {
	return c.detail
}

func (c *Career) StartYear() CareerStartYear {
	return c.startYear
}

func (c *Career) EndYear() CareerEndYear {
	return c.endYear
}
