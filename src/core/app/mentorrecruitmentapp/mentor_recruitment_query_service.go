//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/app/mentorrecruitmentapp/mentor_recruitment_query_service_mock.go -package=mentorrecruitmentapp_mock
package mentorrecruitmentapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
)

type MentorRecruitmentQueryService interface {
	GetMentorRecruitments(ctx context.Context) ([]*mentor_recruitmentdm.MentorRecruitment, error)
}
