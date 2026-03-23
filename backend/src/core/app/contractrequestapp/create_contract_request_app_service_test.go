package contractrequestapp

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	contract_requestdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/contract_requestdm"
	plandm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/plandm"
	userdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/userdm"
)

func TestCreateContractRequestAppService_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockContractRequestRepo := contract_requestdm_mock.NewMockContractRequestRepository(ctrl)
	mockPlanRepo := plandm_mock.NewMockPlanRepository(ctrl)
	mockIsExistByUserID := userdm_mock.NewMockIsExistByUserIDDomainService(ctrl)

	service := NewCreateContractRequestAppService(
		mockContractRequestRepo,
		mockPlanRepo,
		mockIsExistByUserID,
	)

	t.Run("正常系: 契約リクエストが正常に作成される", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()
		planID := shared.NewUUID().String()
		plan := createTestPlan(t, planID, "公開")

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  "プランに興味があります。よろしくお願いします。",
			PlanID:   planID,
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockPlanRepo.EXPECT().
			FindByID(ctx, gomock.Any()).
			Return(plan, nil)

		mockContractRequestRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(nil)

		err := service.Exec(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("異常系: 無効なMenteeID", func(t *testing.T) {
		ctx := context.Background()
		planID := shared.NewUUID().String()

		req := &CreateContractRequestRequest{
			MenteeID: "",
			Message:  "メッセージ",
			PlanID:   planID,
		}

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "UUID")
	})

	t.Run("異常系: 無効なPlanID", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  "メッセージ",
			PlanID:   "invalid-uuid",
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "UUID")
	})

	t.Run("異常系: ユーザー存在チェックでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()
		planID := shared.NewUUID().String()

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  "メッセージ",
			PlanID:   planID,
		}

		expectedError := errors.New("database connection error")

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to check if user exists")
	})

	t.Run("異常系: ユーザーが存在しない", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()
		planID := shared.NewUUID().String()

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  "メッセージ",
			PlanID:   planID,
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(false, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})

	t.Run("異常系: プラン取得でエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()
		planID := shared.NewUUID().String()

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  "メッセージ",
			PlanID:   planID,
		}

		expectedError := errors.New("database error")

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockPlanRepo.EXPECT().
			FindByID(ctx, gomock.Any()).
			Return(nil, expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to find plan")
	})

	t.Run("異常系: プランが存在しない", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()
		planID := shared.NewUUID().String()

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  "メッセージ",
			PlanID:   planID,
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockPlanRepo.EXPECT().
			FindByID(ctx, gomock.Any()).
			Return(nil, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "plan not found")
	})

	t.Run("異常系: プランが中止済み", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()
		planID := shared.NewUUID().String()
		plan := createTestPlan(t, planID, "中止")

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  "メッセージ",
			PlanID:   planID,
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockPlanRepo.EXPECT().
			FindByID(ctx, gomock.Any()).
			Return(plan, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "plan is cancelled")
	})

	t.Run("異常系: メッセージが空でContractRequest作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()
		planID := shared.NewUUID().String()
		plan := createTestPlan(t, planID, "公開")

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  "",
			PlanID:   planID,
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockPlanRepo.EXPECT().
			FindByID(ctx, gomock.Any()).
			Return(plan, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "message must not be empty")
	})

	t.Run("異常系: メッセージが501文字以上でContractRequest作成に失敗", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()
		planID := shared.NewUUID().String()
		plan := createTestPlan(t, planID, "公開")

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  strings.Repeat("あ", 501),
			PlanID:   planID,
		}

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockPlanRepo.EXPECT().
			FindByID(ctx, gomock.Any()).
			Return(plan, nil)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "message must be less than 500 characters")
	})

	t.Run("異常系: リポジトリのStoreでエラーが発生", func(t *testing.T) {
		ctx := context.Background()
		menteeID := shared.NewUUID().String()
		planID := shared.NewUUID().String()
		plan := createTestPlan(t, planID, "公開")

		req := &CreateContractRequestRequest{
			MenteeID: menteeID,
			Message:  "メッセージ",
			PlanID:   planID,
		}

		expectedError := errors.New("repository save error")

		mockIsExistByUserID.EXPECT().
			Exec(ctx, gomock.Any()).
			Return(true, nil)

		mockPlanRepo.EXPECT().
			FindByID(ctx, gomock.Any()).
			Return(plan, nil)

		mockContractRequestRepo.EXPECT().
			Store(ctx, gomock.Any()).
			Return(expectedError)

		err := service.Exec(ctx, req)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func createTestPlan(t *testing.T, planIDStr string, statusStr string) *plandm.Plan {
	t.Helper()

	planID, err := shared.NewUUIDByVal(planIDStr)
	if err != nil {
		t.Fatalf("failed to create plan id: %v", err)
	}

	mentorID := userdm.NewUserID()
	categoryID, err := categorydm.NewCategoryIDByVal(shared.NewUUID().String())
	if err != nil {
		t.Fatalf("failed to create category id: %v", err)
	}

	status, err := plandm.NewStatus(statusStr)
	if err != nil {
		t.Fatalf("failed to create status: %v", err)
	}

	consultationType, err := plandm.NewConsultationType("単発")
	if err != nil {
		t.Fatalf("failed to create consultation type: %v", err)
	}

	plan, err := plandm.NewPlanByVal(
		planID,
		mentorID,
		"テストプラン",
		categoryID,
		[]shared.UUID{},
		"テスト説明",
		status,
		&consultationType,
		5000,
	)
	if err != nil {
		t.Fatalf("failed to create plan: %v", err)
	}

	return plan
}
