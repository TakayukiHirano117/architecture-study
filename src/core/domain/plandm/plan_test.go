package plandm_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func createValidTagIDs(t *testing.T, count int) []tagdm.TagID {
	t.Helper()

	tagIDs := make([]tagdm.TagID, count)
	for i := 0; i < count; i++ {
		tagIDs[i] = tagdm.NewTagID()
	}
	return tagIDs
}

func createValidMentorID(t *testing.T) userdm.UserID {
	t.Helper()
	return userdm.NewUserID()
}

func TestNewPlan(t *testing.T) {
	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32)
		wantErr    bool
		errMsg     string
		assertions func(t *testing.T, plan *plandm.Plan)
	}{
		{
			name: "有効な値でPlanを作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 3), "有効な説明", status, consultationType, 5000
			},
			wantErr: false,
			assertions: func(t *testing.T, plan *plandm.Plan) {
				assert.NotNil(t, plan)
			},
		},
		{
			name: "タイトルが50文字でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), strings.Repeat("あ", 50), categoryID, createValidTagIDs(t, 0), "有効な説明", status, consultationType, 3000
			},
			wantErr: false,
			assertions: func(t *testing.T, plan *plandm.Plan) {
				assert.NotNil(t, plan)
			},
		},
		{
			name: "タグが5個でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 5), "有効な説明", status, consultationType, 5000
			},
			wantErr: false,
			assertions: func(t *testing.T, plan *plandm.Plan) {
				assert.NotNil(t, plan)
			},
		},
		{
			name: "説明が5000文字でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 0), strings.Repeat("あ", 5000), status, consultationType, 5000
			},
			wantErr: false,
			assertions: func(t *testing.T, plan *plandm.Plan) {
				assert.NotNil(t, plan)
			},
		},
		{
			name: "価格が最小値3000でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 0), "有効な説明", status, consultationType, 3000
			},
			wantErr: false,
			assertions: func(t *testing.T, plan *plandm.Plan) {
				assert.NotNil(t, plan)
			},
		},
		{
			name: "価格が最大値1000000でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 0), "有効な説明", status, consultationType, 1000000
			},
			wantErr: false,
			assertions: func(t *testing.T, plan *plandm.Plan) {
				assert.NotNil(t, plan)
			},
		},
		{
			name: "タイトルが空の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "", categoryID, createValidTagIDs(t, 0), "有効な説明", status, consultationType, 5000
			},
			wantErr: true,
			errMsg:  "title must not be empty",
		},
		{
			name: "タイトルが51文字以上の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), strings.Repeat("あ", 51), categoryID, createValidTagIDs(t, 0), "有効な説明", status, consultationType, 5000
			},
			wantErr: true,
			errMsg:  "title must be less than 50 characters",
		},
		{
			name: "タグが6個以上の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 6), "有効な説明", status, consultationType, 5000
			},
			wantErr: true,
			errMsg:  "tags must be less than 5",
		},
		{
			name: "説明が空の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 0), "", status, consultationType, 5000
			},
			wantErr: true,
			errMsg:  "description must not be empty",
		},
		{
			name: "説明が5001文字以上の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 0), strings.Repeat("あ", 5001), status, consultationType, 5000
			},
			wantErr: true,
			errMsg:  "description must be less than 5000 characters",
		},
		{
			name: "価格が3000未満の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 0), "有効な説明", status, consultationType, 2999
			},
			wantErr: true,
			errMsg:  "price must be at least 3000",
		},
		{
			name: "価格が1000000を超える場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 0), "有効な説明", status, consultationType, 1000001
			},
			wantErr: true,
			errMsg:  "price must be less than 1000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, mentorID, title, categoryID, tagIDs, description, status, consultationType, price := tt.setupFunc(t)

			plan, err := plandm.NewPlan(id, mentorID, title, categoryID, tagIDs, description, status, &consultationType, price)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
				return
			}

			require.NoError(t, err)
			if tt.assertions != nil {
				tt.assertions(t, plan)
			}
		})
	}
}

func TestNewPlanByVal(t *testing.T) {
	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32)
		assertions func(t *testing.T, plan *plandm.Plan)
	}{
		{
			name: "DBから取得したデータでPlanを再構築できる",
			setupFunc: func(t *testing.T) (shared.UUID, userdm.UserID, string, categorydm.CategoryID, []tagdm.TagID, string, plandm.Status, plandm.ConsultationType, uint32) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				categoryID, err := categorydm.NewCategoryIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				status, err := plandm.NewStatus("公開")
				require.NoError(t, err)

				consultationType, err := plandm.NewConsultationType("単発")
				require.NoError(t, err)

				return id, createValidMentorID(t), "有効なタイトル", categoryID, createValidTagIDs(t, 3), "有効な説明", status, consultationType, 5000
			},
			assertions: func(t *testing.T, plan *plandm.Plan) {
				assert.NotNil(t, plan)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, mentorID, title, categoryID, tagIDs, description, status, consultationType, price := tt.setupFunc(t)

			plan, err := plandm.NewPlanByVal(id, mentorID, title, categoryID, tagIDs, description, status, &consultationType, price)

			require.NoError(t, err)
			if tt.assertions != nil {
				tt.assertions(t, plan)
			}
		})
	}
}
