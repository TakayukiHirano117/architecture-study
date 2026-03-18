package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	mentorrecruitmentapp_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/app/mentorrecruitmentapp"
)

func createTestMentorRecruitment(t *testing.T) *mentor_recruitmentdm.MentorRecruitment {
	t.Helper()

	id, err := mentor_recruitmentdm.NewMentorRecruitmentIDByVal(shared.NewUUID().String())
	require.NoError(t, err)

	userID := shared.NewUUID()
	categoryID, err := categorydm.NewCategoryIDByVal(shared.NewUUID().String())
	require.NoError(t, err)

	consultationType, err := plandm.NewConsultationTypeByVal("単発")
	require.NoError(t, err)

	consultationMethod, err := mentor_recruitmentdm.NewConsultationMethodByVal("チャット")
	require.NoError(t, err)

	appPeriod, err := mentor_recruitmentdm.NewApplicationPeriodByVal(time.Now().AddDate(0, 0, 14))
	require.NoError(t, err)

	status, err := plandm.NewStatusByVal("公開")
	require.NoError(t, err)

	tagName, err := tagdm.NewTagNameByVal("Go")
	require.NoError(t, err)

	tag, err := tagdm.NewTagByVal(shared.NewUUID(), tagName)
	require.NoError(t, err)

	now := time.Now()
	mr, err := mentor_recruitmentdm.NewMentorRecruitmentByVal(
		id,
		userID,
		"Go言語メンター募集",
		"Go言語を教えてくれるメンターを探しています。",
		categoryID,
		consultationType,
		consultationMethod,
		10000,
		50000,
		appPeriod,
		status,
		[]tagdm.Tag{*tag},
		now,
		now,
	)
	require.NoError(t, err)
	return mr
}

func TestGetMentorRecruitmentsController_Exec(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("正常系: メンター募集一覧が取得でき、data に DTO が含まれる", func(t *testing.T) {
		mockAppService := mentorrecruitmentapp_mock.NewMockGetMentorRecruitmentsAppService(ctrl)
		mr := createTestMentorRecruitment(t)

		controller := &GetMentorRecruitmentsController{
			getMentorRecruitmentsAppService: mockAppService,
		}

		mockAppService.EXPECT().
			Exec(gomock.Any()).
			Return([]*mentor_recruitmentdm.MentorRecruitment{mr}, nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/mentor-recruitments", nil)

		controller.Exec(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Data    []struct {
				ID                 string `json:"id"`
				UserID             string `json:"user_id"`
				Title              string `json:"title"`
				Description        string `json:"description"`
				CategoryID         string `json:"category_id"`
				ConsultationType   string `json:"consultation_type"`
				ConsultationMethod string `json:"consultation_method"`
				BudgetFrom         uint32 `json:"budget_from"`
				BudgetTo           uint32 `json:"budget_to"`
				Status             string `json:"status"`
			} `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "success", response.Status)
		assert.Equal(t, "Mentor recruitments fetched successfully", response.Message)
		require.Len(t, response.Data, 1)
		assert.Equal(t, mr.ID().String(), response.Data[0].ID)
		assert.Equal(t, mr.UserID().String(), response.Data[0].UserID)
		assert.Equal(t, "Go言語メンター募集", response.Data[0].Title)
		assert.Equal(t, "Go言語を教えてくれるメンターを探しています。", response.Data[0].Description)
		assert.Equal(t, mr.CategoryID().String(), response.Data[0].CategoryID)
		assert.Equal(t, string(mr.ConsultationType()), response.Data[0].ConsultationType)
		assert.Equal(t, string(mr.ConsultationMethod()), response.Data[0].ConsultationMethod)
		assert.Equal(t, uint32(10000), response.Data[0].BudgetFrom)
		assert.Equal(t, uint32(50000), response.Data[0].BudgetTo)
		assert.Equal(t, string(mr.Status()), response.Data[0].Status)
	})

	t.Run("異常系: AppService がエラーを返した場合、ctx.Error が呼ばれる", func(t *testing.T) {
		mockAppService := mentorrecruitmentapp_mock.NewMockGetMentorRecruitmentsAppService(ctrl)
		appErr := errors.New("database error")

		controller := &GetMentorRecruitmentsController{
			getMentorRecruitmentsAppService: mockAppService,
		}

		mockAppService.EXPECT().
			Exec(gomock.Any()).
			Return(nil, appErr)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/mentor-recruitments", nil)

		controller.Exec(c)

		require.Len(t, c.Errors, 1)
		assert.Equal(t, appErr.Error(), c.Errors.Last().Err.Error())
	})
}
