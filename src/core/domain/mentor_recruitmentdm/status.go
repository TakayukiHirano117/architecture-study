package mentor_recruitmentdm

import "github.com/cockroachdb/errors"

type Status string

const (
	Published Status = "公開"
	Cancelled Status = "中止"
)

func NewStatus(value string) (Status, error) {
	status := Status(value)
	if status != Published && status != Cancelled {
		return "", errors.New("status must be 公開 or 中止")
	}

	return status, nil
}

func NewStatusByVal(value string) (Status, error) {
	if value == "" {
		return "", errors.New("status must not be empty")
	}

	return Status(value), nil
}

func (s Status) String() (string, error) {
	switch s {
	case Published:
		return string(Published), nil
	case Cancelled:
		return string(Cancelled), nil
	}

	return "", errors.New("invalid status")
}

func (s Status) Equal(s2 Status) bool {
	return s == s2
}
