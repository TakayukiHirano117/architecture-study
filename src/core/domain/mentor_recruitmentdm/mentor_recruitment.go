package mentor_recruitmentdm

import (
	"time"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
)

type MentorRecruitment struct {
	createdAt          time.Time
	updatedAt          time.Time
	id                 MentorRecruitmentID
	title              Title
	description        Description
	category           categorydm.Category
	consultationType   ConsultationType
	consultationMethod ConsultationMethod
	budget             Budget
	applicationPeriod  ApplicationPeriod
	status             Status
	tags               []tagdm.Tag
}

func NewMentorRecruitment(
	id MentorRecruitmentID,
	title Title,
	description Description,
	category categorydm.Category,
	consultationType ConsultationType,
	consultationMethod ConsultationMethod,
	budget Budget,
	applicationPeriod ApplicationPeriod,
	status Status,
	tags []tagdm.Tag,
) (*MentorRecruitment, error) {
	return &MentorRecruitment{
		id:                 id,
		title:              title,
		description:        description,
		category:           category,
		consultationType:   consultationType,
		consultationMethod: consultationMethod,
		budget:             budget,
		applicationPeriod:  applicationPeriod,
		status:             status,
		tags:               tags,
		createdAt:          time.Now(),
		updatedAt:          time.Now(),
	}, nil
}

func NewMentorRecruitmentByVal(
	id MentorRecruitmentID,
	title Title,
	description Description,
	category categorydm.Category,
	consultationType ConsultationType,
	consultationMethod ConsultationMethod,
	budget Budget,
	applicationPeriod ApplicationPeriod,
	status Status,
	tags []tagdm.Tag,
	createdAt time.Time,
	updatedAt time.Time,
) (*MentorRecruitment, error) {
	return &MentorRecruitment{
		id:                 id,
		title:              title,
		description:        description,
		category:           category,
		consultationType:   consultationType,
		consultationMethod: consultationMethod,
		budget:             budget,
		applicationPeriod:  applicationPeriod,
		status:             status,
		tags:               tags,
		createdAt:          createdAt,
		updatedAt:          updatedAt,
	}, nil
}

// TODO: getterやドメインルールを表すメソッドを追加する
