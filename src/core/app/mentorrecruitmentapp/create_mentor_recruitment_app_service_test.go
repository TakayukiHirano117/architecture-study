package mentorrecruitmentapp

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	categorydm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/categorydm"
	mentor_recruitmentdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/mentor_recruitmentdm"
	tagdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/tagdm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"
)

func TestCreateMentorRecruitmentAppService_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIsExistByUserID := userdm_mock.NewMockIsExistByUserIDDomainService(ctrl)
	mockIsExistByCategoryID := categorydm_mock.NewMockIsExistByCategoryIDDomainService(ctrl)
	mockMentorRecruitmentRepo := mentor_recruitmentdm_mock.NewMockMentorRecruitmentRepository(ctrl)
	mockBuildTags := tagdm_mock.NewMockBuildTagsDomainService(ctrl)

	service := NewCreateMentorRecruitmentAppService(
		mockIsExistByUserID,
		mockIsExistByCategoryID,
		mockMentorRecruitmentRepo,
		mockBuildTags,
	)

	t.Run("正常系: タグ付きでメンター募集が正常に作成される", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())
		tagID := uuid.New().String()

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "Goのメンタリング",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "Goの初心者向けメンタリングを募集します。",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags: []CreateMentorRecruitmentTagRequest{
				{ID: tagID, Name: "Go"},
			},
		}

		goTag := createTestTag(t, tagID, "Go")

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{*goTag}, nil)

		mockMentorRecruitmentRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("正常系: タグなしでメンター募集が正常に作成される", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "Pythonのメンタリング",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Continuous,
			ConsultationMethod: mentor_recruitmentdm.VideoCall,
			Description:        "Pythonの中級者向けメンタリングを募集します。",
			BudgetFrom:         10000,
			BudgetTo:           50000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{}, nil)

		mockMentorRecruitmentRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("異常系: 無効なUserID", func(t *testing.T) {
		ctx := context.Background()
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userdm.UserID(""),
			Title:              "テスト",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "UserID is empty")
	})

	t.Run("異常系: ユーザー存在チェックでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テスト",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		expectedError := errors.New("database connection error")

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to check if user exists")
	})

	t.Run("異常系: ユーザーが存在しない", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テスト",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})

	t.Run("異常系: 無効なCategoryID", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テスト",
			CategoryID:         categorydm.CategoryID(""),
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "CategoryID must not be empty")
	})

	t.Run("異常系: カテゴリ存在チェックでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テスト",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		expectedError := errors.New("database connection error")

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to check if category exists")
	})

	t.Run("異常系: カテゴリが存在しない", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テスト",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "category not found")
	})

	t.Run("異常系: タグのビルドでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())
		tagID := uuid.New().String()

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テスト",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags: []CreateMentorRecruitmentTagRequest{
				{ID: tagID, Name: "Go"},
			},
		}

		expectedError := errors.New("some tags not found")

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(nil, expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to build tags")
	})

	t.Run("異常系: タイトルが空でメンター募集作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{}, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "title must not be empty")
	})

	t.Run("異常系: 説明が空でメンター募集作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テストタイトル",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{}, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "description must not be empty")
	})

	t.Run("異常系: 予算下限が最小値未満でメンター募集作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テストタイトル",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         100, // MinBudget(1000)未満
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{}, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "minimum budget must be at least")
	})

	t.Run("異常系: 予算上限が最大値超過でメンター募集作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テストタイトル",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           2000000, // MaxBudget(1000000)超過
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{}, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "maximum budget must be at most")
	})

	t.Run("異常系: 予算下限が上限を超えている場合にメンター募集作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テストタイトル",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         50000,
			BudgetTo:           10000, // BudgetFrom > BudgetTo
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{}, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "minimum budget must be less than or equal to maximum budget")
	})

	t.Run("異常系: リポジトリのStoreでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		userID := userdm.UserID(uuid.New().String())
		categoryID := categorydm.CategoryID(uuid.New().String())

		req := &CreateMentorRecruitmentRequest{
			UserID:             userID,
			Title:              "テストタイトル",
			CategoryID:         categoryID,
			ConsultationType:   plandm.Once,
			ConsultationMethod: mentor_recruitmentdm.Chat,
			Description:        "テスト説明",
			BudgetFrom:         5000,
			BudgetTo:           10000,
			Tags:               []CreateMentorRecruitmentTagRequest{},
		}

		expectedError := errors.New("repository save error")

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{}, nil)

		mockMentorRecruitmentRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

// createTestTag はテスト用のタグを作成するヘルパー関数
func createTestTag(t *testing.T, tagIDStr string, name string) *tagdm.Tag {
	t.Helper()

	tagID, err := shared.NewUUIDByVal(tagIDStr)
	if err != nil {
		t.Fatalf("failed to create tag id: %v", err)
	}

	tagName, err := tagdm.NewTagNameByVal(name)
	if err != nil {
		t.Fatalf("failed to create tag name: %v", err)
	}

	tag, err := tagdm.NewTagByVal(tagID, tagName)
	if err != nil {
		t.Fatalf("failed to create tag: %v", err)
	}

	return tag
}
