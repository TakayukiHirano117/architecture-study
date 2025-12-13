package mentor_recruitmentdm

import (
	"time"
	"unicode/utf8"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
)

type MentorRecruitment struct {
	createdAt          time.Time
	updatedAt          time.Time
	id                 MentorRecruitmentID
	user_id            shared.UUID
	title              string
	description        string
	category_id        categorydm.CategoryID
	consultationType   plandm.ConsultationType
	consultationMethod ConsultationMethod
	budgetFrom         uint32
	budgetTo           uint32
	applicationPeriod  ApplicationPeriod
	status             plandm.Status
	tags               []tagdm.Tag
}

const (
	MinBudget uint32 = 1000
	MaxBudget uint32 = 1000000
)

func NewMentorRecruitment(
	id MentorRecruitmentID,
	user_id shared.UUID,
	title string,
	description string,
	category_id categorydm.CategoryID,
	consultationType plandm.ConsultationType,
	consultationMethod ConsultationMethod,
	budgetFrom uint32,
	budgetTo uint32,
	applicationPeriod ApplicationPeriod,
	status plandm.Status,
	tags []tagdm.Tag,
) (*MentorRecruitment, error) {
	if title == "" {
		return nil, errors.New("title must not be empty")
	}

	if utf8.RuneCountInString(title) > 255 {
		return nil, errors.New("title must be less than 255 characters")
	}

	if description == "" {
		return nil, errors.New("description must not be empty")
	}

	if utf8.RuneCountInString(description) > 2000 {
		return nil, errors.New("description must be less than 2000 characters")
	}

	if budgetFrom < MinBudget {
		return nil, errors.Newf("minimum budget must be at least %d", MinBudget)
	}

	if budgetTo > MaxBudget {
		return nil, errors.Newf("maximum budget must be at most %d", MaxBudget)
	}

	if budgetFrom > budgetTo {
		return nil, errors.New("minimum budget must be less than or equal to maximum budget")
	}

	return &MentorRecruitment{
		id:                 id,
		user_id:            user_id,
		title:              title,
		description:        description,
		category_id:        category_id,
		consultationType:   consultationType,
		consultationMethod: consultationMethod,
		budgetFrom:         budgetFrom,
		budgetTo:           budgetTo,
		applicationPeriod:  applicationPeriod,
		status:             status,
		tags:               tags,
		createdAt:          time.Now(),
		updatedAt:          time.Now(),
	}, nil
}

func NewMentorRecruitmentByVal(
	id MentorRecruitmentID,
	user_id shared.UUID,
	title string,
	description string,
	category_id categorydm.CategoryID,
	consultationType plandm.ConsultationType,
	consultationMethod ConsultationMethod,
	budgetFrom uint32,
	budgetTo uint32,
	applicationPeriod ApplicationPeriod,
	status plandm.Status,
	tags []tagdm.Tag,
	createdAt time.Time,
	updatedAt time.Time,
) (*MentorRecruitment, error) {
	return &MentorRecruitment{
		id:                 id,
		user_id:            user_id,
		title:              title,
		description:        description,
		category_id:        category_id,
		consultationType:   consultationType,
		consultationMethod: consultationMethod,
		budgetFrom:         budgetFrom,
		budgetTo:           budgetTo,
		applicationPeriod:  applicationPeriod,
		status:             status,
		tags:               tags,
		createdAt:          createdAt,
		updatedAt:          updatedAt,
	}, nil
}

func (mr *MentorRecruitment) ID() MentorRecruitmentID {
	return mr.id
}

func (mr *MentorRecruitment) UserID() shared.UUID {
	return mr.user_id
}

func (mr *MentorRecruitment) Title() string {
	return mr.title
}

func (mr *MentorRecruitment) Description() string {
	return mr.description
}

func (mr *MentorRecruitment) CategoryID() categorydm.CategoryID {
	return mr.category_id
}

func (mr *MentorRecruitment) ConsultationType() plandm.ConsultationType {
	return mr.consultationType
}

func (mr *MentorRecruitment) ConsultationMethod() ConsultationMethod {
	return mr.consultationMethod
}

func (mr *MentorRecruitment) BudgetFrom() uint32 {
	return mr.budgetFrom
}

func (mr *MentorRecruitment) BudgetTo() uint32 {
	return mr.budgetTo
}

func (mr *MentorRecruitment) ApplicationPeriod() ApplicationPeriod {
	return mr.applicationPeriod
}

func (mr *MentorRecruitment) Status() plandm.Status {
	return mr.status
}

func (mr *MentorRecruitment) Tags() []tagdm.Tag {
	return mr.tags
}
