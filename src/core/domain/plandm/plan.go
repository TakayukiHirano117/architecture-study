package plandm

import (
	"unicode/utf8"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type Plan struct {
	id               shared.UUID
	mentor_id        userdm.UserID
	title            string
	category_id      categorydm.CategoryID
	tag_ids          []tagdm.TagID
	description      string
	status           Status
	consultationType *ConsultationType
	price            uint32
}

func NewPlan(id shared.UUID, mentor_id userdm.UserID, title string, category_id categorydm.CategoryID, tag_ids []tagdm.TagID, description string, status Status, consultationType *ConsultationType, price uint32) (*Plan, error) {
	if title == "" {
		return nil, errors.New("title must not be empty")
	}

	if utf8.RuneCountInString(title) > 50 {
		return nil, errors.New("title must be less than 50 characters")
	}

	if len(tag_ids) > 5 {
		return nil, errors.New("tags must be less than 5")
	}

	if description == "" {
		return nil, errors.New("description must not be empty")
	}

	if utf8.RuneCountInString(description) > 5000 {
		return nil, errors.New("description must be less than 5000 characters")
	}

	if price < 3000 {
		return nil, errors.New("price must be at least 3000")
	}

	if price > 1000000 {
		return nil, errors.New("price must be less than 1000000")
	}

	return &Plan{id: id, mentor_id: mentor_id, title: title, category_id: category_id, tag_ids: tag_ids, description: description, status: status, consultationType: consultationType, price: price}, nil
}

func NewPlanByVal(id shared.UUID, mentor_id userdm.UserID, title string, category_id categorydm.CategoryID, tag_ids []tagdm.TagID, description string, status Status, consultationType *ConsultationType, price uint32) (*Plan, error) {
	return &Plan{id: id, mentor_id: mentor_id, title: title, category_id: category_id, tag_ids: tag_ids, description: description, status: status, consultationType: consultationType, price: price}, nil
}

func (p *Plan) ID() shared.UUID {
	return p.id
}

func (p *Plan) MentorID() userdm.UserID {
	return p.mentor_id
}

func (p *Plan) Title() string {
	return p.title
}

func (p *Plan) CategoryID() categorydm.CategoryID {
	return p.category_id
}

func (p *Plan) TagIDs() []tagdm.TagID {
	return p.tag_ids
}

func (p *Plan) Description() string {
	return p.description
}

func (p *Plan) Status() Status {
	return p.status
}

func (p *Plan) ConsultationType() *ConsultationType {
	return p.consultationType
}

func (p *Plan) Price() uint32 {
	return p.price
}
