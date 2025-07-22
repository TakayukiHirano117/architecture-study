package userdm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type UserId string

func NewUserId() UserId {
	return UserId(uuid.New().String())
}

func NewUserIdByVal(val string) (UserId, error) {
	if val == "" {
		return "", errors.New("UserId is empty")
	}

	if _, err := uuid.Parse(val); err != nil {
		return "", errors.New("UserId is invalid")
	}

	return UserId(val), nil
}

func (userId UserId) String() string {
	return string(userId)
}

func (userId UserId) Equal(userId2 UserId) bool {
	return userId == userId2
}
