package userapp

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	tagdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/tagdm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"
)

func TestUpdateUserAppService_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := userdm_mock.NewMockUserRepository(ctrl)
	mockIsExistByUserNameExcludeSelf := userdm_mock.NewMockIsExistByUserNameExcludeSelfDomainService(ctrl)
	mockBuildTags := tagdm_mock.NewMockBuildTagsDomainService(ctrl)

	service := NewUpdateUserAppService(
		mockUserRepo,
		mockIsExistByUserNameExcludeSelf,
		mockBuildTags,
	)

	t.Run("正常系: 既存タグIDを使用してユーザーが正常に更新される", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		goTagID := uuid.New().String()
		skillID := uuid.New().String()
		careerID := uuid.New().String()

		req := &UpdateUserRequest{
			ID:               userID,
			Name:             "updated_user",
			Email:            "updated@example.com",
			SelfIntroduction: "Updated introduction",
			Skills: []UpdateSkillRequest{
				{
					ID: skillID,
					Tag: UpdateTagRequest{
						ID:   goTagID,
						Name: "Go",
					},
					Evaluation:        5,
					YearsOfExperience: 4,
				},
			},
			Careers: []UpdateCareerRequest{
				{
					ID:        careerID,
					Detail:    "Updated career detail",
					StartYear: 2020,
					EndYear:   2024,
				},
			},
		}

		// 既存ユーザーを作成
		existingUser := createTestUser(t, userID)

		// BuildTags が返すタグを作成
		goTag := createTestTag(t, goTagID, "Go")

		mockIsExistByUserNameExcludeSelf.EXPECT().
			Exec(ctx, gomock.Any(), gomock.Any()).
			Return(false, nil)

		mockUserRepo.EXPECT().
			FindByID(ctx, gomock.Any()).
			Return(existingUser, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{*goTag}, nil)

		mockUserRepo.EXPECT().
			Update(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("正常系: 新規タグを使用してユーザーが正常に更新される", func(t *testing.T) {
		ctx := context.Background()
		userID := uuid.New().String()
		newTagID := uuid.New().String()

		req := &UpdateUserRequest{
			ID:               userID,
			Name:             "updated_user",
			Email:            "updated@example.com",
			SelfIntroduction: "Updated introduction",
			Skills: []UpdateSkillRequest{
				{
					ID: "",
					Tag: UpdateTagRequest{
						ID:   "",
						Name: "NewTag",
					},
					Evaluation:        4,
					YearsOfExperience: 2,
				},
			},
			Careers: []UpdateCareerRequest{
				{
					ID:        "",
					Detail:    "New career detail",
					StartYear: 2021,
					EndYear:   2024,
				},
			},
		}

		existingUser := createTestUser(t, userID)

		// BuildTags が返す新規タグを作成
		newTag := createTestTag(t, newTagID, "NewTag")

		mockIsExistByUserNameExcludeSelf.EXPECT().
			Exec(ctx, gomock.Any(), gomock.Any()).
			Return(false, nil)

		mockUserRepo.EXPECT().
			FindByID(ctx, gomock.Any()).
			Return(existingUser, nil)

		mockBuildTags.EXPECT().
			Exec(ctx, gomock.Any()).
			Return([]tagdm.Tag{*newTag}, nil)

		mockUserRepo.EXPECT().
			Update(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
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

// createTestUser はテスト用のユーザーを作成するヘルパー関数
func createTestUser(t *testing.T, userIDStr string) *userdm.User {
	t.Helper()

	userID, err := shared.NewUUIDByVal(userIDStr)
	if err != nil {
		t.Fatalf("failed to create user id: %v", err)
	}

	userName, err := userdm.NewUserName("test_user")
	if err != nil {
		t.Fatalf("failed to create user name: %v", err)
	}

	password, err := userdm.NewPasswordByVal("hashedpassword123")
	if err != nil {
		t.Fatalf("failed to create password: %v", err)
	}

	email, err := userdm.NewEmail("test@example.com")
	if err != nil {
		t.Fatalf("failed to create email: %v", err)
	}

	tagID := shared.NewUUID()
	tagName, err := tagdm.NewTagNameByVal("Go")
	if err != nil {
		t.Fatalf("failed to create tag name: %v", err)
	}

	tag, err := tagdm.NewTagByVal(tagID, tagName)
	if err != nil {
		t.Fatalf("failed to create tag: %v", err)
	}

	evaluation, err := userdm.NewEvaluationByVal(4)
	if err != nil {
		t.Fatalf("failed to create evaluation: %v", err)
	}

	yearsOfExperience, err := userdm.NewYearsOfExperienceByVal(3)
	if err != nil {
		t.Fatalf("failed to create years of experience: %v", err)
	}

	skill, err := userdm.NewSkill(userdm.NewSkillID(), tag, evaluation, yearsOfExperience)
	if err != nil {
		t.Fatalf("failed to create skill: %v", err)
	}

	careerDetail, err := userdm.NewCareerDetail("Software Engineer")
	if err != nil {
		t.Fatalf("failed to create career detail: %v", err)
	}

	startYear, err := userdm.NewCareerStartYear(2020)
	if err != nil {
		t.Fatalf("failed to create start year: %v", err)
	}

	endYear, err := userdm.NewCareerEndYear(2023)
	if err != nil {
		t.Fatalf("failed to create end year: %v", err)
	}

	career, err := userdm.NewCareer(userdm.NewCareerID(), *careerDetail, *startYear, *endYear)
	if err != nil {
		t.Fatalf("failed to create career: %v", err)
	}

	selfIntroduction, err := userdm.NewSelfIntroduction("Test introduction")
	if err != nil {
		t.Fatalf("failed to create self introduction: %v", err)
	}

	user, err := userdm.NewUserByVal(userID, *userName, password, *email, []userdm.Skill{*skill}, []userdm.Career{*career}, selfIntroduction)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	return user
}
