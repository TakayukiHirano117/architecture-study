package mentor_recruitmentdm_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
)

type mentorRecruitmentParams struct {
	id                 mentor_recruitmentdm.MentorRecruitmentID
	title              string
	description        string
	category           categorydm.Category
	consultationType   mentor_recruitmentdm.ConsultationType
	consultationMethod mentor_recruitmentdm.ConsultationMethod
	budgetFrom         uint32
	budgetTo           uint32
	applicationPeriod  mentor_recruitmentdm.ApplicationPeriod
	status             mentor_recruitmentdm.Status
	tags               []tagdm.Tag
}

func createValidMentorRecruitmentParams(t *testing.T) mentorRecruitmentParams {
	t.Helper()

	consultationType, err := mentor_recruitmentdm.NewConsultationType("単発")
	require.NoError(t, err)

	consultationMethod, err := mentor_recruitmentdm.NewConsultationMethod("チャット")
	require.NoError(t, err)

	status, err := mentor_recruitmentdm.NewStatus("公開")
	require.NoError(t, err)

	return mentorRecruitmentParams{
		id:                 mentor_recruitmentdm.NewMentorRecruitmentID(),
		title:              "Goのメンターを募集します",
		description:        "Go言語の学習をサポートしてくれるメンターを探しています。",
		category:           createValidCategory(t),
		consultationType:   consultationType,
		consultationMethod: consultationMethod,
		budgetFrom:         5000,
		budgetTo:           10000,
		applicationPeriod:  mentor_recruitmentdm.NewApplicationPeriod(),
		status:             status,
		tags:               []tagdm.Tag{createValidTag(t)},
	}
}

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
	t.Run("有効なパラメータでMentorRecruitmentを作成できる", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		require.NoError(t, err)
		assert.NotNil(t, mr)
	})

	t.Run("tagsが空でもMentorRecruitmentを作成できる", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.tags = []tagdm.Tag{}

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		require.NoError(t, err)
		assert.NotNil(t, mr)
	})
}

func TestMentorRecruitment_NewMentorRecruitment_TitleValidation(t *testing.T) {
	t.Run("タイトルが空文字の場合はエラー", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.title = ""

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		assert.Error(t, err)
		assert.Nil(t, mr)
	})

	t.Run("タイトルが255文字の場合は作成できる", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.title = strings.Repeat("あ", 255)

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		require.NoError(t, err)
		assert.NotNil(t, mr)
	})

	t.Run("タイトルが256文字以上の場合はエラー", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.title = strings.Repeat("あ", 256)

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		assert.Error(t, err)
		assert.Nil(t, mr)
	})
}

func TestMentorRecruitment_NewMentorRecruitment_DescriptionValidation(t *testing.T) {
	t.Run("説明が空文字の場合はエラー", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.description = ""

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		assert.Error(t, err)
		assert.Nil(t, mr)
	})

	t.Run("説明が2000文字の場合は作成できる", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.description = strings.Repeat("あ", 2000)

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		require.NoError(t, err)
		assert.NotNil(t, mr)
	})

	t.Run("説明が2001文字以上の場合はエラー", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.description = strings.Repeat("あ", 2001)

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		assert.Error(t, err)
		assert.Nil(t, mr)
	})
}

func TestMentorRecruitment_NewMentorRecruitment_BudgetValidation(t *testing.T) {
	t.Run("budgetFromが最小値(1000)の場合は作成できる", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.budgetFrom = 1000
		cvmrp.budgetTo = 10000

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		require.NoError(t, err)
		assert.NotNil(t, mr)
	})

	t.Run("budgetFromが最小値未満の場合はエラー", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.budgetFrom = 999
		cvmrp.budgetTo = 10000

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		assert.Error(t, err)
		assert.Nil(t, mr)
	})

	t.Run("budgetToが最大値(1000000)の場合は作成できる", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.budgetFrom = 5000
		cvmrp.budgetTo = 1000000

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		require.NoError(t, err)
		assert.NotNil(t, mr)
	})

	t.Run("budgetToが最大値を超える場合はエラー", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.budgetFrom = 5000
		cvmrp.budgetTo = 1000001

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		assert.Error(t, err)
		assert.Nil(t, mr)
	})

	t.Run("budgetFromがbudgetToより大きい場合はエラー", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.budgetFrom = 20000
		cvmrp.budgetTo = 10000

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		assert.Error(t, err)
		assert.Nil(t, mr)
	})

	t.Run("budgetFromとbudgetToが同じ値の場合は作成できる", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		cvmrp.budgetFrom = 5000
		cvmrp.budgetTo = 5000

		mr, err := mentor_recruitmentdm.NewMentorRecruitment(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
		)

		require.NoError(t, err)
		assert.NotNil(t, mr)
	})
}

func TestMentorRecruitment_NewMentorRecruitmentByVal(t *testing.T) {
	t.Run("有効なパラメータでMentorRecruitmentを作成できる", func(t *testing.T) {
		cvmrp := createValidMentorRecruitmentParams(t)
		createdAt := time.Now()
		updatedAt := time.Now()

		mr, err := mentor_recruitmentdm.NewMentorRecruitmentByVal(
			cvmrp.id,
			cvmrp.title,
			cvmrp.description,
			cvmrp.category,
			cvmrp.consultationType,
			cvmrp.consultationMethod,
			cvmrp.budgetFrom,
			cvmrp.budgetTo,
			cvmrp.applicationPeriod,
			cvmrp.status,
			cvmrp.tags,
			createdAt,
			updatedAt,
		)

		require.NoError(t, err)
		assert.NotNil(t, mr)
	})
}
