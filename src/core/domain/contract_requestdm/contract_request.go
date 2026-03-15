package contract_requestdm

import (
	"unicode/utf8"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

type ContractRequest struct {
	id             shared.UUID
	message        string
	menteeID       shared.UUID
	priceAtRequest uint32
	planID         shared.UUID
	isAccepted     IsAccepted
}

func NewContractRequest(id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted IsAccepted) (*ContractRequest, error) {
	if message == "" {
		return nil, errors.New("message must not be empty")
	}

	if utf8.RuneCountInString(message) > 500 {
		return nil, errors.New("message must be less than 500 characters")
	}

	if priceAtRequest < 3000 {
		return nil, errors.New("price must be at least 3000")
	}

	if priceAtRequest > 1000000 {
		return nil, errors.New("priceAtRequest must be less than 1000000")
	}

	return &ContractRequest{id: id, message: message, menteeID: menteeID, priceAtRequest: priceAtRequest, planID: planID, isAccepted: isAccepted}, nil
}

func NewContractRequestByVal(id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted IsAccepted) (*ContractRequest, error) {
	return &ContractRequest{id: id, message: message, menteeID: menteeID, priceAtRequest: priceAtRequest, planID: planID, isAccepted: isAccepted}, nil
}

func (c *ContractRequest) ID() shared.UUID {
	return c.id
}

func (c *ContractRequest) Message() string {
	return c.message
}

func (c *ContractRequest) MenteeID() shared.UUID {
	return c.menteeID
}

func (c *ContractRequest) PriceAtRequest() uint32 {
	return c.priceAtRequest
}

func (c *ContractRequest) PlanID() shared.UUID {
	return c.planID
}

func (c *ContractRequest) IsAccepted() IsAccepted {
	return c.isAccepted
}
