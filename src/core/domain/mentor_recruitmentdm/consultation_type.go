package mentor_recruitmentdm

import "github.com/cockroachdb/errors"

type ConsultationType string

const (
	Once       ConsultationType = "単発"
	Continuous ConsultationType = "継続"
)

func NewConsultationType(value string) (ConsultationType, error) {
	t := ConsultationType(value)
	if t != Once && t != Continuous {
		return "", errors.New("consultation type must be 単発 or 継続")
	}

	return t, nil
}

func NewConsultationTypeByVal(value string) (ConsultationType, error) {
	if value == "" {
		return "", errors.New("consultation type must not be empty")
	}

	return ConsultationType(value), nil
}

func (c ConsultationType) String() (string, error) {
	switch c {
	case Once:
		return string(Once), nil
	case Continuous:
		return string(Continuous), nil
	}

	return "", errors.New("invalid consultation type")
}

func (c ConsultationType) Equal(c2 ConsultationType) bool {
	return c == c2
}
