package shared

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type UUID string

func NewUUID() UUID {
	return UUID(uuid.New().String())
}

func NewUUIDByVal(val string) (UUID, error) {
	if val == "" {
		return "", errors.New("UUID is empty")
	}

	if _, err := uuid.Parse(val); err != nil {
		return "", errors.New("UUID is invalid")
	}

	return UUID(val), nil
}

func (uuid UUID) String() string {
	return string(uuid)
}

func (uuid UUID) Equal(uuid2 UUID) bool {
	return uuid == uuid2
}
