package mentor_recruitmentdm

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type MentorRecruitmentID string

func NewMentorRecruitmentID() MentorRecruitmentID {
	return MentorRecruitmentID(uuid.New().String())
}

func NewMentorRecruitmentIDByVal(val string) (MentorRecruitmentID, error) {
	if val == "" {
		return "", errors.New("MentorRecruitmentID must not be empty")
	}

	if _, err := uuid.Parse(val); err != nil {
		return "", errors.New("MentorRecruitmentID must be a valid UUID")
	}

	return MentorRecruitmentID(val), nil
}

func (mentorRecruitmentId MentorRecruitmentID) String() string {
	return string(mentorRecruitmentId)
}

func (mentorRecruitmentId MentorRecruitmentID) Equal(mentorRecruitmentId2 MentorRecruitmentID) bool {
	return mentorRecruitmentId == mentorRecruitmentId2
}
