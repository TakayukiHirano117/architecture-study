package userapp

import (
	"context"
	"errors"
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAppService_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockUserDomainService := mocks.NewMockUserDomainService(ctrl)
	mockTagDomainService := mocks.NewMockTagDomainService(ctrl)

	service := NewCreateUserAppService(mockUserRepo, mockUserDomainService, mockTagDomainService)

	t.Run("正常系: ユーザーが正常に作成される", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					TagName:           "Go",
					Evaluation:        4,
					YearsOfExperience: 3,
				},
				{
					TagName:           "Python",
					Evaluation:        3,
					YearsOfExperience: 2,
				},
			},
			Careers: []CreateCareerRequest{
				{
					Detail:    "Software Engineer at ABC Corp",
					StartYear: 2020,
					EndYear:   2023,
				},
			},
			SelfIntroduction: "I am a passionate software engineer.",
		}

		// Mock設定: ユーザー名の重複チェック（存在しない）
		mockUserDomainService.EXPECT().
			IsExistByUserName(ctx, gomock.Any()).
			Return(false, nil)

		// Mock設定: タグの存在チェック（Go）
		goTagId, _ := tagdm.NewTagId(uuid.New().String())
		mockTagDomainService.EXPECT().
			IsExistByTagName(ctx, gomock.Any()).
			Return(&goTagId, nil)

		// Mock設定: タグの存在チェック（Python）
		pythonTagId, _ := tagdm.NewTagId(uuid.New().String())
		mockTagDomainService.EXPECT().
			IsExistByTagName(ctx, gomock.Any()).
			Return(&pythonTagId, nil)

		// Mock設定: ユーザーの保存
		mockUserRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(nil)

		// 実行
		err := service.Exec(ctx, req)

		// 検証
		assert.NoError(t, err)
	})

	t.Run("異常系: ユーザー名が既に存在する", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "existing_user",
			Email:            "test@example.com",
			Password:         "password123",
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		// Mock設定: ユーザー名の重複チェック（存在する）
		mockUserDomainService.EXPECT().
			IsExistByUserName(ctx, gomock.Any()).
			Return(true, nil)

		// 実行
		err := service.Exec(ctx, req)

		// 検証
		assert.Error(t, err)
		assert.Equal(t, "user name already exists", err.Error())
	})

	t.Run("異常系: タグが存在しない", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					TagName:           "NonExistentTag",
					Evaluation:        4,
					YearsOfExperience: 3,
				},
			},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		// Mock設定: ユーザー名の重複チェック（存在しない）
		mockUserDomainService.EXPECT().
			IsExistByUserName(ctx, gomock.Any()).
			Return(false, nil)

		// Mock設定: タグの存在チェック（存在しない）
		mockTagDomainService.EXPECT().
			IsExistByTagName(ctx, gomock.Any()).
			Return(nil, nil)

		// 実行
		err := service.Exec(ctx, req)

		// 検証
		assert.Error(t, err)
		assert.Equal(t, "tag name not found", err.Error())
	})

	t.Run("異常系: ユーザー名チェックでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "test_user",
			Email:            "test@example.com",
			Password:         "password123",
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		expectedError := errors.New("database connection error")

		// Mock設定: ユーザー名の重複チェックでエラー
		mockUserDomainService.EXPECT().
			IsExistByUserName(ctx, gomock.Any()).
			Return(false, expectedError)

		// 実行
		err := service.Exec(ctx, req)

		// 検証
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("異常系: タグ存在チェックでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					TagName:           "Go",
					Evaluation:        4,
					YearsOfExperience: 3,
				},
			},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		expectedError := errors.New("tag service error")

		// Mock設定: ユーザー名の重複チェック（存在しない）
		mockUserDomainService.EXPECT().
			IsExistByUserName(ctx, gomock.Any()).
			Return(false, nil)

		// Mock設定: タグの存在チェックでエラー
		mockTagDomainService.EXPECT().
			IsExistByTagName(ctx, gomock.Any()).
			Return(nil, expectedError)

		// 実行
		err := service.Exec(ctx, req)

		// 検証
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("異常系: ユーザー保存でエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					TagName:           "Go",
					Evaluation:        4,
					YearsOfExperience: 3,
				},
			},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		expectedError := errors.New("repository save error")

		// Mock設定: ユーザー名の重複チェック（存在しない）
		mockUserDomainService.EXPECT().
			IsExistByUserName(ctx, gomock.Any()).
			Return(false, nil)

		// Mock設定: タグの存在チェック（Go）
		goTagId, _ := tagdm.NewTagId(uuid.New().String())
		mockTagDomainService.EXPECT().
			IsExistByTagName(ctx, gomock.Any()).
			Return(&goTagId, nil)

		// Mock設定: ユーザーの保存でエラー
		mockUserRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(expectedError)

		// 実行
		err := service.Exec(ctx, req)

		// 検証
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestCreateUserAppService_Exec_InvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockUserDomainService := mocks.NewMockUserDomainService(ctrl)
	mockTagDomainService := mocks.NewMockTagDomainService(ctrl)

	service := NewCreateUserAppService(mockUserRepo, mockUserDomainService, mockTagDomainService)

	t.Run("異常系: 無効なユーザー名", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "", // 空のユーザー名
			Email:            "test@example.com",
			Password:         "password123456",
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		// 実行
		err := service.Exec(ctx, req)

		// 検証
		assert.Error(t, err)
	})

	t.Run("異常系: 無効なメールアドレス", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "test_user",
			Email:            "invalid-email", // 無効なメールアドレス
			Password:         "password123456",
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		// Mock設定: ユーザー名の重複チェック（存在しない）
		mockUserDomainService.EXPECT().
			IsExistByUserName(ctx, gomock.Any()).
			Return(false, nil)

		// 実行
		err := service.Exec(ctx, req)

		// 検証
		assert.Error(t, err)
	})

	t.Run("異常系: 無効なパスワード", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "test_user",
			Email:            "test@example.com",
			Password:         "short", // 短すぎるパスワード
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		// Mock設定: ユーザー名の重複チェック（存在しない）
		mockUserDomainService.EXPECT().
			IsExistByUserName(ctx, gomock.Any()).
			Return(false, nil)

		// 実行
		err := service.Exec(ctx, req)

		// 検証
		assert.Error(t, err)
	})
}
