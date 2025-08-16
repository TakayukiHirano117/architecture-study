package userdm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type CareerID string

func NewCareerID() CareerID {
	return CareerID(uuid.New().String())
}

func NewCareerIDByVal(val string) (CareerID, error) {
	if val == "" {
		return "", errors.New("CareerID is empty")
	}

	return CareerID(val), nil
}

func (careerId CareerID) String() string {
	return string(careerId)
}

func (careerId CareerID) Equal(careerId2 CareerID) bool {
	return careerId == careerId2
}
