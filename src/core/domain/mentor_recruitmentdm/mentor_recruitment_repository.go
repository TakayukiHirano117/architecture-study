//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/mentor_recruitmentdm/mentor_recruitment_repository_mock.go -package=mentor_recruitmentdm_mock
package mentor_recruitmentdm

import "context"

type MentorRecruitmentRepository interface {
	Store(ctx context.Context, mentorRecruitment *MentorRecruitment) error
}