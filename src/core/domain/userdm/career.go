package userdm

type Career struct {
	careerId  CareerId
	detail    string
	startYear int
	endYear   int
}

func NewCareer(careerId CareerId, detail string, startYear int, endYear int) (*Career, error) {
	return &Career{
		careerId:  careerId,
		detail:    detail,
		startYear: startYear,
		endYear:   endYear,
	}, nil
}
