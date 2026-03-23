package contract_requestdm

import "github.com/cockroachdb/errors"

type IsAccepted string

const (
	Accepted IsAccepted = "承認"
	Pending  IsAccepted = "未確認"
	Rejected IsAccepted = "拒否"
)

func NewIsAccepted(value string) (IsAccepted, error) {
	isAccepted := IsAccepted(value)

	if isAccepted == "" {
		return "", errors.New("is accepted must not be empty")
	}

	if isAccepted != Accepted && isAccepted != Pending && isAccepted != Rejected {
		return "", errors.New("is accepted must be 承認 or 未確認 or 拒否")
	}

	return isAccepted, nil
}

func NewIsAcceptedByVal(value string) (IsAccepted, error) {
	if value == "" {
		return "", errors.New("is accepted must not be empty")
	}

	return IsAccepted(value), nil
}

func (i IsAccepted) String() string {
	return string(i)
}

func (i IsAccepted) Equal(i2 IsAccepted) bool {
	return i == i2
}
