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
	tagdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/tagdm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"
)

func TestCreateUserController_Exec(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("正常系: ユーザーが正常に作成される", func(t *testing.T) {
		mockUserRepo := userdm_mock.NewMockUserRepository(ctrl)
		mockIsExistByUserName := userdm_mock.NewMockIsExistByUserNameDomainService(ctrl)
		mockBuildTags := tagdm_mock.NewMockBuildTagsDomainService(ctrl)

		controller := &CreateUserController{
			userRepo:          mockUserRepo,
			IsExistByUserName: mockIsExistByUserName,
			BuildTags:         mockBuildTags,
		}

		tagID := uuid.New().String()
		reqBody := map[string]interface{}{
			"name":              "test_user",
			"email":             "test@example.com",
			"password":          "password123456",
			"self_introduction": "Hello, I am a test user.",
			"skills": []map[string]interface{}{
				{
					"tag": map[string]interface{}{
						"id":   tagID,
						"name": "Go",
					},
					"evaluation":          4,
					"years_of_experience": 3,
				},
			},
			"careers": []map[string]interface{}{
				{
					"detail":     "Software Engineer",
					"start_year": 2020,
					"end_year":   2023,
				},
			},
		}
		body, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		// モックの設定
		goTag := createControllerTestTag(t, tagID, "Go")

		mockIsExistByUserName.EXPECT().
			Exec(gomock.Any(), gomock.Any()).
			Return(false, nil)

		mockBuildTags.EXPECT().
			Exec(gomock.Any(), gomock.Any()).
			Return([]tagdm.Tag{*goTag}, nil)

		mockUserRepo.EXPECT().
			Store(gomock.Any(), gomock.Any()).
			Return(nil)

		controller.Exec(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])
	})

	t.Run("異常系: 不正なJSONでBadRequest", func(t *testing.T) {
		controller := &CreateUserController{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString("invalid json"))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.Exec(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

// createControllerTestTag はテスト用のタグを作成するヘルパー関数
func createControllerTestTag(t *testing.T, tagIDStr string, name string) *tagdm.Tag {
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

