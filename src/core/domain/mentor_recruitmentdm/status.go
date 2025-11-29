package mentor_recruitmentdm

import "github.com/cockroachdb/errors"

type Status string

const (
	published Status = "公開"
	cancelled Status = "中止"
)

func NewStatus(value string) (Status, error) {
	status := Status(value)
	if status != published && status != cancelled {
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
	case published:
		return "公開", nil
	case cancelled:
		return "中止", nil
	}

	return "", errors.New("invalid status")
}

func (s Status) Equal(s2 Status) bool {
	return s == s2
}
