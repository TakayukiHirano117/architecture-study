package contractdm

import (
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/cockroachdb/errors"
	"unicode/utf8"
)

type Contract struct {
	id shared.UUID
	mentee_id shared.UUID
	plan_id shared.UUID
	message string
}

func NewContract(id shared.UUID, mentee_id shared.UUID, plan_id shared.UUID, message string) (*Contract, error) {
	if message == "" {
		return nil, errors.New("message must not be empty")
	}

	if utf8.RuneCountInString(message) > 500 {
		return nil, errors.New("message must be less than 500 characters")
	}

	return &Contract{id: id, mentee_id: mentee_id, plan_id: plan_id, message: message}, nil
}


func NewContractByVal(id shared.UUID, mentee_id shared.UUID, plan_id shared.UUID, message string) (*Contract, error) {
	return &Contract{id: id, mentee_id: mentee_id, plan_id: plan_id, message: message}, nil
}

func (c *Contract) ID() shared.UUID {
	return c.id
}

func (c *Contract) MenteeID() shared.UUID {
	return c.mentee_id
}

func (c *Contract) PlanID() shared.UUID {
	return c.plan_id
}

func (c *Contract) Message() string {
	return c.message
}
