package userdm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func createValidSkill(t *testing.T) *userdm.Skill {
	t.Helper()

	tagId := shared.NewUUID()
	tagName, err := tagdm.NewTagName("test-tag-name")
	require.NoError(t, err)

	tag, err := tagdm.NewTag(tagId, *tagName)
	require.NoError(t, err)

	skill, err := userdm.NewSkill(userdm.NewSkillID(), tag, 5, 3)
	require.NoError(t, err)

	return skill
}

func createValidCareer(t *testing.T) *userdm.Career {
	t.Helper()

	careerDetail, err := userdm.NewCareerDetail("Web開発に従事")
	require.NoError(t, err)

	careerStartYear, err := userdm.NewCareerStartYear(2020)
	require.NoError(t, err)

	careerEndYear, err := userdm.NewCareerEndYear(2022)
	require.NoError(t, err)

	career, err := userdm.NewCareer(
		userdm.NewCareerID(),
		*careerDetail,
		*careerStartYear,
		*careerEndYear,
	)
	require.NoError(t, err)

	return career
}

func TestUser_NewUser(t *testing.T) {
	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) (shared.UUID, userdm.UserName, userdm.Password, userdm.Email, []userdm.Skill, []userdm.Career, *userdm.SelfIntroduction)
		wantErr    bool
		assertions func(t *testing.T, user *userdm.User)
	}{
		{
			name: "有効なパラメータでUserを作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserName, userdm.Password, userdm.Email, []userdm.Skill, []userdm.Career, *userdm.SelfIntroduction) {
				userId := shared.NewUUID()

				userName, err := userdm.NewUserName("Test User")
				require.NoError(t, err)

				password, err := userdm.NewPassword("validPassword0")
				require.NoError(t, err)

				email, err := userdm.NewEmail("test@example.com")
				require.NoError(t, err)

				selfIntroduction, err := userdm.NewSelfIntroduction("よろしくお願いします")
				require.NoError(t, err)

				skill := createValidSkill(t)
				skills := []userdm.Skill{*skill}

				career := createValidCareer(t)
				careers := []userdm.Career{*career}

				return userId, *userName, *password, *email, skills, careers, selfIntroduction
			},
			wantErr: false,
			assertions: func(t *testing.T, user *userdm.User) {
				assert.NotNil(t, user)
				assert.Equal(t, 1, len(user.Skills()))
				assert.Equal(t, 1, len(user.Careers()))
			},
		},
		{
			name: "空のスキルリストはエラー",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserName, userdm.Password, userdm.Email, []userdm.Skill, []userdm.Career, *userdm.SelfIntroduction) {
				userId := shared.NewUUID()

				userName, err := userdm.NewUserName("Test User")
				require.NoError(t, err)

				password, err := userdm.NewPassword("validPassword0")
				require.NoError(t, err)

				email, err := userdm.NewEmail("test@example.com")
				require.NoError(t, err)

				selfIntroduction, err := userdm.NewSelfIntroduction("よろしくお願いします")
				require.NoError(t, err)

				skills := []userdm.Skill{}
				careers := []userdm.Career{}

				return userId, *userName, *password, *email, skills, careers, selfIntroduction
			},
			wantErr:    true,
			assertions: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId, userName, password, email, skills, careers, selfIntroduction := tt.setupFunc(t)

			user, err := userdm.NewUser(userId, userName, password, email, skills, careers, selfIntroduction)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			if tt.assertions != nil {
				tt.assertions(t, user)
			}
		})
	}
}
