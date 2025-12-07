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

	categorydm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/categorydm"
	mentor_recruitmentdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/mentor_recruitmentdm"
	tagdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/tagdm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"
)

func TestCreateMentorRecruitmentController_Exec(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("正常系: メンター募集が正常に作成される", func(t *testing.T) {
		mockMentorRecruitmentRepo := mentor_recruitmentdm_mock.NewMockMentorRecruitmentRepository(ctrl)
		mockBuildTags := tagdm_mock.NewMockBuildTagsDomainService(ctrl)
		mockUserRepo := userdm_mock.NewMockUserRepository(ctrl)
		mockIsExistByCategoryID := categorydm_mock.NewMockIsExistByCategoryIDDomainService(ctrl)
		mockIsExistByUserID := userdm_mock.NewMockIsExistByUserIDDomainService(ctrl)

		controller := &CreateMentorRecruitmentController{
			mentorRecruitmentRepo: mockMentorRecruitmentRepo,
			buildTags:             mockBuildTags,
			userRepo:              mockUserRepo,
			isExistByCategoryID:   mockIsExistByCategoryID,
			isExistByUserID:       mockIsExistByUserID,
		}

		userID := uuid.New().String()
		categoryID := uuid.New().String()

		reqBody := map[string]interface{}{
			"user_id":             userID,
			"title":               "Go言語メンター募集",
			"description":         "Go言語を教えてくれるメンターを探しています。バックエンド開発経験のある方を希望します。",
			"category_id":         categoryID,
			"consultation_type":   "one_time",
			"consultation_method": "online",
			"budget_from":         10000,
			"budget_to":           50000,
			"tags": []map[string]interface{}{
				{
					"id":   uuid.New().String(),
					"name": "Go",
				},
			},
		}
		body, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/mentor_recruitments", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		// モックの設定
		mockIsExistByUserID.EXPECT().
			Exec(gomock.Any(), gomock.Any()).
			Return(true, nil)

		mockIsExistByCategoryID.EXPECT().
			Exec(gomock.Any(), gomock.Any()).
			Return(true, nil)

		mockBuildTags.EXPECT().
			Exec(gomock.Any(), gomock.Any()).
			Return(nil, nil)

		mockMentorRecruitmentRepo.EXPECT().
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
		controller := &CreateMentorRecruitmentController{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/mentor_recruitments", bytes.NewBufferString("invalid json"))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.Exec(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
