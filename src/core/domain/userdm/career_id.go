package userdm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type CareerId string

func NewCareerId() CareerId {
	return CareerId(uuid.New().String())
}

func NewCareerIdByVal(val string) (CareerId, error) {
	if val == "" {
		return "", errors.New("CareerId is empty")
	}

	return CareerId(val), nil
}

func (careerId CareerId) String() string {
	return string(careerId)
}

func (careerId CareerId) Equal(careerId2 CareerId) bool {
	return careerId == careerId2
}
