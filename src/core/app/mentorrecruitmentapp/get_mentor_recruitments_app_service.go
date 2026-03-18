//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/app/mentorrecruitmentapp/get_mentor_recruitments_app_service_mock.go -package=mentorrecruitmentapp_mock
package mentorrecruitmentapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
)

// GetMentorRecruitmentsAppService はメンター募集一覧取得のアプリケーションサービスインターフェース
type GetMentorRecruitmentsAppService interface {
	Exec(ctx context.Context) ([]*mentor_recruitmentdm.MentorRecruitment, error)
}

type getMentorRecruitmentsAppServiceImpl struct {
	mentorRecruitmentQueryService MentorRecruitmentQueryService
}

func NewGetMentorRecruitmentsAppService() GetMentorRecruitmentsAppService {
	return &getMentorRecruitmentsAppServiceImpl{
		mentorRecruitmentQueryService: rdbimpl.NewMentorRecruitmentQueryServiceImpl(),
	}
}

func (s *getMentorRecruitmentsAppServiceImpl) Exec(ctx context.Context) ([]*mentor_recruitmentdm.MentorRecruitment, error) {
	mentorRecruitments, err := s.mentorRecruitmentQueryService.GetMentorRecruitments(ctx)
	if err != nil {
		return nil, err
	}

	return mentorRecruitments, nil
}
