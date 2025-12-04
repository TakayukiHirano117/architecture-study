package plandm

import (
	"unicode/utf8"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/vo"
)

type Plan struct {
	id               vo.UUID
	title            string
	category_id      categorydm.CategoryID
	tags             []tagdm.Tag
	content          string
	status           vo.Status
	consultationType *vo.ConsultationType
	price            uint32
}

func NewPlan(id vo.UUID, title string, category_id categorydm.CategoryID, tags []tagdm.Tag, content string, status vo.Status, consultationType vo.ConsultationType, price uint32) (*Plan, error) {
	if title == "" {
		return nil, errors.New("title must not be empty")
	}

	if utf8.RuneCountInString(title) > 50 {
		return nil, errors.New("title must be less than 50 characters")
	}

	if len(tags) > 5 {
		return nil, errors.New("tags must be less than 5")
	}

	if content == "" {
		return nil, errors.New("content must not be empty")
	}

	if utf8.RuneCountInString(content) > 5000 {
		return nil, errors.New("content must be less than 5000 characters")
	}

	if price < 3000 {
		return nil, errors.New("price must be at least 3000")
	}

	if price > 1000000 {
		return nil, errors.New("price must be less than 1000000")
	}

	return &Plan{id: id, title: title, category_id: category_id, tags: tags, content: content, status: status, consultationType: &consultationType, price: price}, nil
}

func NewPlanByVal(id vo.UUID, title string, category_id categorydm.CategoryID, tags []tagdm.Tag, content string, status vo.Status, consultationType vo.ConsultationType, price uint32) (*Plan, error) {
	return &Plan{id: id, title: title, category_id: category_id, tags: tags, content: content, status: status, consultationType: &consultationType, price: price}, nil
}

// Getter
