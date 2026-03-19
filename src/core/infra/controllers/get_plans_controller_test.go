package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	planapp_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/app/planapp"
)

func createTestPlan(t *testing.T) *plandm.Plan {
	t.Helper()

	id := shared.NewUUID()
	mentorID, err := userdm.NewUserIDByVal(shared.NewUUID().String())
	require.NoError(t, err)

	categoryID, err := categorydm.NewCategoryIDByVal(shared.NewUUID().String())
	require.NoError(t, err)

	status, err := plandm.NewStatus("公開")
	require.NoError(t, err)

	consultationType, err := plandm.NewConsultationType("単発")
	require.NoError(t, err)

	tagIDs := []shared.UUID{shared.NewUUID()}

	plan, err := plandm.NewPlanByVal(
		id,
		mentorID,
		"Goメンタリングプラン",
		categoryID,
		tagIDs,
		"Go言語の基礎から応用までサポートします。",
		status,
		&consultationType,
		10000,
	)
	require.NoError(t, err)
	return plan
}

func TestGetPlansController_Exec(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("正常系: プラン一覧が取得でき、data に DTO が含まれる", func(t *testing.T) {
		mockAppService := planapp_mock.NewMockGetPlansAppService(ctrl)
		plan := createTestPlan(t)

		controller := &GetPlansController{
			getPlansAppService: mockAppService,
		}

		mockAppService.EXPECT().
			Exec(gomock.Any()).
			Return([]*plandm.Plan{plan}, nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/plans", nil)

		controller.Exec(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Data    []struct {
				ID               string   `json:"id"`
				MentorID         string   `json:"mentor_id"`
				Title            string   `json:"title"`
				CategoryID       string   `json:"category_id"`
				Description      string   `json:"description"`
				Status           string   `json:"status"`
				ConsultationType string   `json:"consultation_type"`
				Price            uint32   `json:"price"`
				TagIDs           []string `json:"tag_ids"`
			} `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "success", response.Status)
		assert.Equal(t, "Plans fetched successfully", response.Message)
		require.Len(t, response.Data, 1)
		assert.Equal(t, plan.ID().String(), response.Data[0].ID)
		assert.Equal(t, plan.MentorID().String(), response.Data[0].MentorID)
		assert.Equal(t, "Goメンタリングプラン", response.Data[0].Title)
		assert.Equal(t, plan.CategoryID().String(), response.Data[0].CategoryID)
		assert.Equal(t, "Go言語の基礎から応用までサポートします。", response.Data[0].Description)
		assert.Equal(t, string(plan.Status()), response.Data[0].Status)
		assert.Equal(t, string(*plan.ConsultationType()), response.Data[0].ConsultationType)
		assert.Equal(t, uint32(10000), response.Data[0].Price)
		require.Len(t, response.Data[0].TagIDs, 1)
		assert.Equal(t, plan.TagIDs()[0].String(), response.Data[0].TagIDs[0])
	})

	t.Run("異常系: AppService がエラーを返した場合、ctx.Error が呼ばれる", func(t *testing.T) {
		mockAppService := planapp_mock.NewMockGetPlansAppService(ctrl)
		appErr := errors.New("database error")

		controller := &GetPlansController{
			getPlansAppService: mockAppService,
		}

		mockAppService.EXPECT().
			Exec(gomock.Any()).
			Return(nil, appErr)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/plans", nil)

		controller.Exec(c)

		require.Len(t, c.Errors, 1)
		assert.Equal(t, appErr.Error(), c.Errors.Last().Err.Error())
	})
}
