package planapp

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	categorydm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/categorydm"
	plandm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/plandm"
	tagdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/tagdm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"
)

func TestCreatePlanAppService_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIsExistByUserID := userdm_mock.NewMockIsExistByUserIDDomainService(ctrl)
	mockIsExistByCategoryID := categorydm_mock.NewMockIsExistByCategoryIDDomainService(ctrl)
	mockPlanRepo := plandm_mock.NewMockPlanRepository(ctrl)
	mockBuildTags := tagdm_mock.NewMockBuildTagsDomainService(ctrl)

	service := NewCreatePlanAppService(
		mockIsExistByUserID,
		mockIsExistByCategoryID,
		mockPlanRepo,
		mockBuildTags,
	)

	t.Run("正常系: タグ付きでプランが正常に作成される", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()
		tagID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "Goメンタリングプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{{ID: tagID, Name: "Go"}},
			Content:          "Go言語の基礎から応用までサポートします。",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
		}

		goTag := createPlanTestTag(t, tagID, "Go")

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{*goTag}, nil)

		mockPlanRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("正常系: タグなしでプランが正常に作成される", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "Pythonメンタリングプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "Python言語の基礎から応用までサポートします。",
			Status:           "公開",
			ConsultationType: "継続",
			Price:            50000,
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

		mockPlanRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("異常系: 無効なUserID", func(t *testing.T) {
		ctx := context.Background()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           "",
			Title:            "テストプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
		}

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "UUID is empty")
	})

	t.Run("異常系: ユーザー存在チェックでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
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
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
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
		userID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストプラン",
			CategoryID:       "",
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
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
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
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
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
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
		userID := uuid.New().String()
		categoryID := uuid.New().String()
		tagID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{{ID: tagID, Name: "Go"}},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
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

	t.Run("異常系: 無効なStatus", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "invalid_status",
			ConsultationType: "単発",
			Price:            10000,
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
		assert.Contains(t, err.Error(), "status must be 公開 or 中止")
	})

	t.Run("異常系: 無効なConsultationType", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストプラン",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "invalid_type",
			Price:            10000,
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
		assert.Contains(t, err.Error(), "consultation type must be 単発 or 継続")
	})

	t.Run("異常系: タイトルが空でプラン作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
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

	t.Run("異常系: 説明が空でプラン作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストタイトル",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
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

	t.Run("異常系: 価格が最小値未満でプラン作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストタイトル",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            1000, // 3000未満
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
		assert.Contains(t, err.Error(), "price must be at least 3000")
	})

	t.Run("異常系: 価格が最大値超過でプラン作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストタイトル",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            2000000, // 1000000超過
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
		assert.Contains(t, err.Error(), "price must be less than 1000000")
	})

	t.Run("異常系: タグが5個を超えてプラン作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		tagIDs := make([]string, 6)
		tags := make([]CreatePlanTagRequest, 6)
		tagObjects := make([]tagdm.Tag, 6)
		for i := 0; i < 6; i++ {
			tagIDs[i] = uuid.New().String()
			tags[i] = CreatePlanTagRequest{ID: tagIDs[i], Name: "Tag"}
			tagObjects[i] = *createPlanTestTag(t, tagIDs[i], "Tag")
		}

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストタイトル",
			CategoryID:       categoryID,
			Tags:             tags,
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(tagObjects, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "tags must be less than 5")
	})

	t.Run("異常系: リポジトリのStoreでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		req := &CreatePlanRequest{
			UserID:           userID,
			Title:            "テストタイトル",
			CategoryID:       categoryID,
			Tags:             []CreatePlanTagRequest{},
			Content:          "テスト説明",
			Status:           "公開",
			ConsultationType: "単発",
			Price:            10000,
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

		mockPlanRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func createPlanTestTag(t *testing.T, tagIDStr string, name string) *tagdm.Tag {
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
