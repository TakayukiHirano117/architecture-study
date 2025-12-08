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
	plandm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/plandm"
	tagdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/tagdm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"
)

func TestCreatePlanController_Exec(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("正常系: プランが正常に作成される", func(t *testing.T) {
		mockPlanRepo := plandm_mock.NewMockPlanRepository(ctrl)
		mockBuildTags := tagdm_mock.NewMockBuildTagsDomainService(ctrl)
		mockIsExistByCategoryID := categorydm_mock.NewMockIsExistByCategoryIDDomainService(ctrl)
		mockIsExistByUserID := userdm_mock.NewMockIsExistByUserIDDomainService(ctrl)

		controller := &CreatePlanController{
			planRepo:            mockPlanRepo,
			buildTags:           mockBuildTags,
			isExistByCategoryID: mockIsExistByCategoryID,
			isExistByUserID:     mockIsExistByUserID,
		}

		userID := uuid.New().String()
		categoryID := uuid.New().String()

		reqBody := map[string]interface{}{
			"user_id":           userID,
			"title":             "Goメンタリングプラン",
			"category_id":       categoryID,
			"content":           "Go言語の基礎から応用までサポートします。バックエンド開発経験のある方を希望します。",
			"status":            "公開",
			"consultation_type": "単発",
			"price":             10000,
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
		c.Request = httptest.NewRequest(http.MethodPost, "/plans", bytes.NewBuffer(body))
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

		mockPlanRepo.EXPECT().
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
		controller := &CreatePlanController{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/plans", bytes.NewBufferString("invalid json"))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.Exec(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
