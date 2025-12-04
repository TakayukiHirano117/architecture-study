package userapp

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	tagdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/tagdm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"
)

func TestCreateUserAppService_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userdm_mock.NewMockUserRepository(ctrl)
	mockIsExistByUserName := userdm_mock.NewMockIsExistByUserNameDomainService(ctrl)
	mockBuildTags := tagdm_mock.NewMockBuildTagsDomainService(ctrl)

	service := NewCreateUserAppService(mockUserRepo, mockIsExistByUserName, mockBuildTags)

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

		// BuildTags が返すタグを作成
		goTag := createTestTagForCreate(t, goTagID, "Go")
		pythonTag := createTestTagForCreate(t, pythonTagID, "Python")

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{*goTag, *pythonTag}, nil)

		mockUserRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("正常系: 新規タグを使用してユーザーが正常に作成される", func(t *testing.T) {
		ctx := context.Background()
		newTagID := uuid.New().String()

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

		// BuildTags が返す新規タグを作成
		newTag := createTestTagForCreate(t, newTagID, "NewTag")

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{*newTag}, nil)

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

	t.Run("異常系: タグIDが存在しない（BuildTags がエラーを返す）", func(t *testing.T) {
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

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(nil, errors.New("some tags not found"))

		err := service.Exec(ctx, req)

		assert.Error(t, err)
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

	t.Run("異常系: BuildTags でエラーが発生", func(t *testing.T) {
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

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(nil, expectedError)

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

		goTag := createTestTagForCreate(t, tagID, "Go")

		mockIsExistByUserName.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{*goTag}, nil)

		mockUserRepo.EXPECT().
			Store(ctx, gomock.Any()).
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
	mockIsExistByUserName := userdm_mock.NewMockIsExistByUserNameDomainService(ctrl)
	mockBuildTags := tagdm_mock.NewMockBuildTagsDomainService(ctrl)

	service := NewCreateUserAppService(mockUserRepo, mockIsExistByUserName, mockBuildTags)

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

// createTestTagForCreate はテスト用のタグを作成するヘルパー関数
func createTestTagForCreate(t *testing.T, tagIDStr string, name string) *tagdm.Tag {
	t.Helper()

	tagID, err := tagdm.NewTagIDByVal(tagIDStr)
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
