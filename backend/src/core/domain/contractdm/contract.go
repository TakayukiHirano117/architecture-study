package contractdm

import (
	"unicode/utf8"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

type Contract struct {
	id       shared.UUID
	menteeID shared.UUID
	planID   shared.UUID
	message  string
}

func NewContract(id shared.UUID, menteeID shared.UUID, planID shared.UUID, message string) (*Contract, error) {
	if message == "" {
		return nil, errors.New("message must not be empty")
	}

	if utf8.RuneCountInString(message) > 500 {
		return nil, errors.New("message must be less than 500 characters")
	}

	return &Contract{id: id, menteeID: menteeID, planID: planID, message: message}, nil
}

func NewContractByVal(id shared.UUID, menteeID shared.UUID, planID shared.UUID, message string) (*Contract, error) {
	return &Contract{id: id, menteeID: menteeID, planID: planID, message: message}, nil
}

func (c *Contract) ID() shared.UUID {
	return c.id
}

func (c *Contract) MenteeID() shared.UUID {
	return c.menteeID
}

func (c *Contract) PlanID() shared.UUID {
	return c.planID
}

func (c *Contract) Message() string {
	return c.message
}
