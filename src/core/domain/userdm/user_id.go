package userdm

import "github.com/google/uuid"

type UserId string

func NewUserId() UserId {
	return UserId(uuid.New().String())
}

func NewUserIdByVal()

func (userId UserId) String() string {
	return string(userId)
}

func (userId UserId) Equal(userId2 UserId) bool {
	return userId == userId2
}
