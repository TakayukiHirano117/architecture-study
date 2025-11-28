package userapp

import (
	"context"
	"errors"
	"testing"

	tagdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/tagdm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserAppService_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userdm_mock.NewMockUserRepository(ctrl)
	mockTagRepo := tagdm_mock.NewMockTagRepository(ctrl)
	mockIsExistByUserName := userdm_mock.NewMockIsExistByUserNameDomainService(ctrl)
	mockIsExistByTagID := tagdm_mock.NewMockIsExistByTagIDDomainService(ctrl)

	service := NewCreateUserAppService(mockUserRepo, mockTagRepo, mockIsExistByUserName, mockIsExistByTagID)

	t.Run("正常系: 既存タグIDを使用してユーザーが正常に作成される", func(t *testing.T) {
		ctx := context.Background()
		goTagID := uuid.New().String()
		pythonTagID := uuid.New().String()

		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					Tag: TagParamRequest{
						ID:   &goTagID,
						Name: "Go",
					},
					Evaluation:        4,
					YearsOfExperience: 3,
				},
				{
					Tag: TagParamRequest{
						ID:   &pythonTagID,
						Name: "Python",
					},
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

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		mockIsExistByTagID.EXPECT().
			Exec(ctx, goTagID).
			Return(true, nil)

		mockIsExistByTagID.EXPECT().
			Exec(ctx, pythonTagID).
			Return(true, nil)

		mockUserRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("正常系: 新規タグを使用してユーザーが正常に作成される", func(t *testing.T) {
		ctx := context.Background()

		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					Tag: TagParamRequest{
						ID:   nil,
						Name: "NewTag",
					},
					Evaluation:        4,
					YearsOfExperience: 3,
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

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		mockTagRepo.EXPECT().
			BulkInsert(ctx, gomock.Any()).
			Return(nil)

		mockUserRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("異常系: ユーザー名が既に存在する", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "existing_user",
			Email:            "test@example.com",
			Password:         "password123456",
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, "user name already exists", err.Error())
	})

	t.Run("異常系: タグIDが存在しない", func(t *testing.T) {
		ctx := context.Background()
		nonExistentTagID := uuid.New().String()

		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					Tag: TagParamRequest{
						ID:   &nonExistentTagID,
						Name: "NonExistentTag",
					},
					Evaluation:        4,
					YearsOfExperience: 3,
				},
			},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		mockIsExistByTagID.EXPECT().
			Exec(ctx, nonExistentTagID).
			Return(false, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "tag with ID")
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("異常系: ユーザー名チェックでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "test_user",
			Email:            "test@example.com",
			Password:         "password123456",
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		expectedError := errors.New("database connection error")

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("異常系: タグ存在確認でエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		tagID := uuid.New().String()

		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					Tag: TagParamRequest{
						ID:   &tagID,
						Name: "Go",
					},
					Evaluation:        4,
					YearsOfExperience: 3,
				},
			},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		expectedError := errors.New("tag service error")

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		mockIsExistByTagID.EXPECT().
			Exec(ctx, tagID).
			Return(false, expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("異常系: ユーザー保存でエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		tagID := uuid.New().String()

		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					Tag: TagParamRequest{
						ID:   &tagID,
						Name: "Go",
					},
					Evaluation:        4,
					YearsOfExperience: 3,
				},
			},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		expectedError := errors.New("repository save error")

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		mockIsExistByTagID.EXPECT().
			Exec(ctx, tagID).
			Return(true, nil)

		mockUserRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("異常系: 新規タグ保存でエラーが発生", func(t *testing.T) {
		ctx := context.Background()

		req := &CreateUserRequest{
			Name:     "test_user",
			Email:    "test@example.com",
			Password: "password123456",
			Skills: []CreateSkillRequest{
				{
					Tag: TagParamRequest{
						ID:   nil,
						Name: "NewTag",
					},
					Evaluation:        4,
					YearsOfExperience: 3,
				},
			},
			Careers: []CreateCareerRequest{
				{
					Detail:    "Software Engineer at ABC Corp",
					StartYear: 2020,
					EndYear:   2023,
				},
			},
			SelfIntroduction: "Test introduction",
		}

		expectedError := errors.New("tag bulk insert error")

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		mockTagRepo.EXPECT().
			BulkInsert(ctx, gomock.Any()).
			Return(expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestCreateUserAppService_Exec_InvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userdm_mock.NewMockUserRepository(ctrl)
	mockTagRepo := tagdm_mock.NewMockTagRepository(ctrl)
	mockIsExistByUserName := userdm_mock.NewMockIsExistByUserNameDomainService(ctrl)
	mockIsExistByTagID := tagdm_mock.NewMockIsExistByTagIDDomainService(ctrl)

	service := NewCreateUserAppService(mockUserRepo, mockTagRepo, mockIsExistByUserName, mockIsExistByTagID)

	t.Run("異常系: 無効なユーザー名", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "",
			Email:            "test@example.com",
			Password:         "password123456",
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		err := service.Exec(ctx, req)

		assert.Error(t, err)
	})

	t.Run("異常系: 無効なメールアドレス", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "test_user",
			Email:            "invalid-email",
			Password:         "password123456",
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
	})

	t.Run("異常系: 無効なパスワード", func(t *testing.T) {
		ctx := context.Background()
		req := &CreateUserRequest{
			Name:             "test_user",
			Email:            "test@example.com",
			Password:         "short",
			Skills:           []CreateSkillRequest{},
			Careers:          []CreateCareerRequest{},
			SelfIntroduction: "Test introduction",
		}

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
	})
}
