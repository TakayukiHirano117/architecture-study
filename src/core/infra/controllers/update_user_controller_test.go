package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	tagdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/tagdm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"
)

func TestUpdateUserController_Exec(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("正常系: ユーザーが正常に更新される", func(t *testing.T) {
		mockUserRepo := userdm_mock.NewMockUserRepository(ctrl)
		mockIsExistByUserName := userdm_mock.NewMockIsExistByUserNameExcludeSelfDomainService(ctrl)
		mockBuildTags := tagdm_mock.NewMockBuildTagsDomainService(ctrl)

		controller := &UpdateUserController{
			userRepo:          mockUserRepo,
			IsExistByUserName: mockIsExistByUserName,
			BuildTags:         mockBuildTags,
		}

		userID := uuid.New().String()
		tagID := uuid.New().String()
		skillID := uuid.New().String()
		careerID := uuid.New().String()

		reqBody := map[string]interface{}{
			"name":              "updated_user",
			"email":             "updated@example.com",
			"self_introduction": "Updated introduction",
			"skills": []map[string]interface{}{
				{
					"id": skillID,
					"tag": map[string]interface{}{
						"id":   tagID,
						"name": "Go",
					},
					"evaluation":          5,
					"years_of_experience": 4,
				},
			},
			"careers": []map[string]interface{}{
				{
					"id":         careerID,
					"detail":     "Updated career",
					"start_year": 2020,
					"end_year":   2024,
				},
			},
		}
		body, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPut, "/users/"+userID, bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Params = gin.Params{{Key: "id", Value: userID}}

		// モックの設定
		existingUser := createUpdateTestUser(t, userID)
		goTag := createControllerTestTag(t, tagID, "Go")

		mockIsExistByUserName.EXPECT().
			Exec(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(false, nil)

		mockUserRepo.EXPECT().
			FindByID(gomock.Any(), gomock.Any()).
			Return(existingUser, nil)

		mockBuildTags.EXPECT().
			Exec(gomock.Any(), gomock.Any()).
			Return([]tagdm.Tag{*goTag}, nil)

		mockUserRepo.EXPECT().
			Update(gomock.Any(), gomock.Any()).
			Return(nil)

		controller.Exec(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])
	})

	t.Run("異常系: 不正なJSONでBadRequest", func(t *testing.T) {
		controller := &UpdateUserController{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPut, "/users/123", bytes.NewBufferString("invalid json"))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Params = gin.Params{{Key: "id", Value: "123"}}

		controller.Exec(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

// createUpdateTestUser はテスト用のユーザーを作成するヘルパー関数
func createUpdateTestUser(t *testing.T, userIDStr string) *userdm.User {
	t.Helper()

	userID, err := userdm.NewUserIDByVal(userIDStr)
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

	tagID := tagdm.NewTagID()
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
