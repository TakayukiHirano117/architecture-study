package mentor_recruitmentdm_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
)

func createValidCategory(t *testing.T) categorydm.Category {
	t.Helper()

	categoryId := categorydm.NewCategoryID()
	categoryName, err := categorydm.NewCategoryName("プログラミング")
	require.NoError(t, err)

	category, err := categorydm.NewCategory(categoryId, *categoryName)
	require.NoError(t, err)

	return *category
}

func createValidTag(t *testing.T) tagdm.Tag {
	t.Helper()

	tagId := tagdm.NewTagID()
	tagName, err := tagdm.NewTagName("Go")
	require.NoError(t, err)

	tag, err := tagdm.NewTag(tagId, *tagName)
	require.NoError(t, err)

	return *tag
}

func TestMentorRecruitment_NewMentorRecruitment(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(t *testing.T) (
			mentor_recruitmentdm.MentorRecruitmentID,
			string,
			string,
			categorydm.Category,
			mentor_recruitmentdm.ConsultationType,
			mentor_recruitmentdm.ConsultationMethod,
			uint32,
			uint32,
			mentor_recruitmentdm.ApplicationPeriod,
			mentor_recruitmentdm.Status,
			[]tagdm.Tag,
		)
		wantErr    bool
		assertions func(t *testing.T, mr *mentor_recruitmentdm.MentorRecruitment)
	}{
		{
			name: "有効なパラメータでMentorRecruitmentを作成できる",
			setupFunc: func(t *testing.T) (
				mentor_recruitmentdm.MentorRecruitmentID,
				string,
				string,
				categorydm.Category,
				mentor_recruitmentdm.ConsultationType,
				mentor_recruitmentdm.ConsultationMethod,
				uint32,
				uint32,
				mentor_recruitmentdm.ApplicationPeriod,
				mentor_recruitmentdm.Status,
				[]tagdm.Tag,
			) {
				id := mentor_recruitmentdm.NewMentorRecruitmentID()

				title := "Goのメンターを募集します"
				description := "Go言語の学習をサポートしてくれるメンターを探しています。"

				category := createValidCategory(t)

				consultationType, err := mentor_recruitmentdm.NewConsultationType("単発")
				require.NoError(t, err)

				consultationMethod, err := mentor_recruitmentdm.NewConsultationMethod("チャット")
				require.NoError(t, err)

				budgetFrom := uint32(5000)
				budgetTo := uint32(10000)
				require.NoError(t, err)

				applicationPeriod := mentor_recruitmentdm.NewApplicationPeriod()

				status, err := mentor_recruitmentdm.NewStatus("公開")
				require.NoError(t, err)

				tags := []tagdm.Tag{createValidTag(t)}

				return id, title, description, category, consultationType, consultationMethod, budgetFrom, budgetTo, applicationPeriod, status, tags
			},
			wantErr: false,
			assertions: func(t *testing.T, mr *mentor_recruitmentdm.MentorRecruitment) {
				assert.NotNil(t, mr)
			},
		},
		{
			name: "tagsが空でもMentorRecruitmentを作成できる（任意項目）",
			setupFunc: func(t *testing.T) (
				mentor_recruitmentdm.MentorRecruitmentID,
				string,
				string,
				categorydm.Category,
				mentor_recruitmentdm.ConsultationType,
				mentor_recruitmentdm.ConsultationMethod,
				uint32,
				uint32,
				mentor_recruitmentdm.ApplicationPeriod,
				mentor_recruitmentdm.Status,
				[]tagdm.Tag,
			) {
				id := mentor_recruitmentdm.NewMentorRecruitmentID()

				title := "Goのメンターを募集します"

				description := "Go言語の学習をサポートしてくれるメンターを探しています。"

				category := createValidCategory(t)

				consultationType, err := mentor_recruitmentdm.NewConsultationType("継続")
				require.NoError(t, err)

				consultationMethod, err := mentor_recruitmentdm.NewConsultationMethod("ビデオ通話")
				require.NoError(t, err)

				budgetFrom := uint32(10000)
				budgetTo := uint32(50000)
				require.NoError(t, err)

				applicationPeriod := mentor_recruitmentdm.NewApplicationPeriod()

				status, err := mentor_recruitmentdm.NewStatus("公開")
				require.NoError(t, err)

				tags := []tagdm.Tag{}

				return id, title, description, category, consultationType, consultationMethod, budgetFrom, budgetTo, applicationPeriod, status, tags
			},
			wantErr: false,
			assertions: func(t *testing.T, mr *mentor_recruitmentdm.MentorRecruitment) {
				assert.NotNil(t, mr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, title, description, category, consultationType, consultationMethod, budgetFrom, budgetTo, applicationPeriod, status, tags := tt.setupFunc(t)

			mr, err := mentor_recruitmentdm.NewMentorRecruitment(
				id,
				title,
				description,
				category,
				consultationType,
				consultationMethod,
				budgetFrom,
				budgetTo,
				applicationPeriod,
				status,
				tags,
			)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			if tt.assertions != nil {
				tt.assertions(t, mr)
			}
		})
	}
}

func TestMentorRecruitment_NewMentorRecruitmentByVal(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(t *testing.T) (
			mentor_recruitmentdm.MentorRecruitmentID,
			string,
			string,
			categorydm.Category,
			mentor_recruitmentdm.ConsultationType,
			mentor_recruitmentdm.ConsultationMethod,
			uint32,
			uint32,
			mentor_recruitmentdm.ApplicationPeriod,
			mentor_recruitmentdm.Status,
			[]tagdm.Tag,
			time.Time,
			time.Time,
		)
		wantErr    bool
		assertions func(t *testing.T, mr *mentor_recruitmentdm.MentorRecruitment)
	}{
		{
			name: "有効なパラメータでMentorRecruitmentを作成できる",
			setupFunc: func(t *testing.T) (
				mentor_recruitmentdm.MentorRecruitmentID,
				string,
				string,
				categorydm.Category,
				mentor_recruitmentdm.ConsultationType,
				mentor_recruitmentdm.ConsultationMethod,
				uint32,
				uint32,
				mentor_recruitmentdm.ApplicationPeriod,
				mentor_recruitmentdm.Status,
				[]tagdm.Tag,
				time.Time,
				time.Time,
			) {
				id := mentor_recruitmentdm.NewMentorRecruitmentID()

				title := "Goのメンターを募集します"

				description := "Go言語の学習をサポートしてくれるメンターを探しています。"

				category := createValidCategory(t)

				consultationType, err := mentor_recruitmentdm.NewConsultationTypeByVal("単発")
				require.NoError(t, err)

				consultationMethod, err := mentor_recruitmentdm.NewConsultationMethodByVal("チャット")
				require.NoError(t, err)

				budgetFrom := uint32(5000)
				budgetTo := uint32(10000)
				require.NoError(t, err)

				applicationPeriod, err := mentor_recruitmentdm.NewApplicationPeriodByVal(time.Now().AddDate(0, 0, 7))
				require.NoError(t, err)

				status, err := mentor_recruitmentdm.NewStatusByVal("公開")
				require.NoError(t, err)

				tags := []tagdm.Tag{createValidTag(t)}
				createdAt := time.Now()
				updatedAt := time.Now()

				return id, title, description, category, consultationType, consultationMethod, budgetFrom, budgetTo, applicationPeriod, status, tags, createdAt, updatedAt
			},
			wantErr: false,
			assertions: func(t *testing.T, mr *mentor_recruitmentdm.MentorRecruitment) {
				assert.NotNil(t, mr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, title, description, category, consultationType, consultationMethod, budgetFrom, budgetTo, applicationPeriod, status, tags, createdAt, updatedAt := tt.setupFunc(t)

			mr, err := mentor_recruitmentdm.NewMentorRecruitmentByVal(
				id,
				title,
				description,
				category,
				consultationType,
				consultationMethod,
				budgetFrom,
				budgetTo,
				applicationPeriod,
				status,
				tags,
				createdAt,
				updatedAt,
			)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			if tt.assertions != nil {
				tt.assertions(t, mr)
			}
		})
	}
}
