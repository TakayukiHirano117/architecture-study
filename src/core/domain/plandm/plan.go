package plandm

import (
	"unicode/utf8"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
)

type Plan struct {
	id               shared.UUID
	title            string
	category_id      categorydm.CategoryID
	tagIDs           []tagdm.TagID
	content          string
	status           Status
	consultationType *ConsultationType
	price            uint32
}

func NewPlan(id shared.UUID, title string, category_id categorydm.CategoryID, tagIDs []tagdm.TagID, content string, status Status, consultationType *ConsultationType, price uint32) (*Plan, error) {
	if title == "" {
		return nil, errors.New("title must not be empty")
	}

	if utf8.RuneCountInString(title) > 50 {
		return nil, errors.New("title must be less than 50 characters")
	}

	if len(tagIDs) > 5 {
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

	return &Plan{id: id, title: title, category_id: category_id, tagIDs: tagIDs, content: content, status: status, consultationType: consultationType, price: price}, nil
}

func NewPlanByVal(id shared.UUID, title string, category_id categorydm.CategoryID, tagIDs []tagdm.TagID, content string, status Status, consultationType *ConsultationType, price uint32) (*Plan, error) {
	return &Plan{id: id, title: title, category_id: category_id, tagIDs: tagIDs, content: content, status: status, consultationType: consultationType, price: price}, nil
}

// Getter
