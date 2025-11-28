package userdm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func NewUserIDByVal(val string) (UserID, error) {
	if val == "" {
		return "", errors.New("UserID is empty")
	}

	if _, err := uuid.Parse(val); err != nil {
		return "", errors.New("UserID is invalid")
	}

	return UserID(val), nil
}

func (userId UserID) String() string {
	return string(userId)
}

func (userId UserID) Equal(userId2 UserID) bool {
	return userId == userId2
}
