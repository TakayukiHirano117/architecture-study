package contract_requestdm_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	contract_requestdm "github.com/TakayukiHirano117/architecture-study/src/core/domain/contract_requestdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

func TestContractRequest_NewContractRequest(t *testing.T) {
	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted)
		wantErr    bool
		errMsg     string
		assertions func(t *testing.T, cr *contract_requestdm.ContractRequest, id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted contract_requestdm.IsAccepted)
	}{
		{
			name: "有効な値でContractRequestを作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, "プランに興味があります。よろしくお願いします。", menteeID, 5000, planID, contract_requestdm.Pending
			},
			wantErr: false,
			assertions: func(t *testing.T, cr *contract_requestdm.ContractRequest, id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted contract_requestdm.IsAccepted) {
				assert.NotNil(t, cr)
				assert.Equal(t, id, cr.ID())
				assert.Equal(t, message, cr.Message())
				assert.Equal(t, menteeID, cr.MenteeID())
				assert.Equal(t, priceAtRequest, cr.PriceAtRequest())
				assert.Equal(t, planID, cr.PlanID())
				assert.Equal(t, isAccepted, cr.IsAccepted())
			},
		},
		{
			name: "messageが500文字でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, strings.Repeat("あ", 500), menteeID, 5000, planID, contract_requestdm.Pending
			},
			wantErr: false,
			assertions: func(t *testing.T, cr *contract_requestdm.ContractRequest, id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted contract_requestdm.IsAccepted) {
				assert.NotNil(t, cr)
				assert.Equal(t, message, cr.Message())
			},
		},
		{
			name: "messageが1文字でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, "a", menteeID, 5000, planID, contract_requestdm.Pending
			},
			wantErr: false,
			assertions: func(t *testing.T, cr *contract_requestdm.ContractRequest, id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted contract_requestdm.IsAccepted) {
				assert.NotNil(t, cr)
			},
		},
		{
			name: "priceAtRequestが3000でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, "メッセージ", menteeID, 3000, planID, contract_requestdm.Pending
			},
			wantErr: false,
			assertions: func(t *testing.T, cr *contract_requestdm.ContractRequest, id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted contract_requestdm.IsAccepted) {
				assert.NotNil(t, cr)
				assert.Equal(t, priceAtRequest, cr.PriceAtRequest())
			},
		},
		{
			name: "priceAtRequestが1000000でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, "メッセージ", menteeID, 1000000, planID, contract_requestdm.Pending
			},
			wantErr: false,
			assertions: func(t *testing.T, cr *contract_requestdm.ContractRequest, id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted contract_requestdm.IsAccepted) {
				assert.NotNil(t, cr)
				assert.Equal(t, priceAtRequest, cr.PriceAtRequest())
			},
		},
		{
			name: "messageが空の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, "", menteeID, 5000, planID, contract_requestdm.Pending
			},
			wantErr: true,
			errMsg:  "message must not be empty",
		},
		{
			name: "messageが501文字以上の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, strings.Repeat("あ", 501), menteeID, 5000, planID, contract_requestdm.Pending
			},
			wantErr: true,
			errMsg:  "message must be less than 500 characters",
		},
		{
			name: "priceAtRequestが3000未満の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, "メッセージ", menteeID, 2999, planID, contract_requestdm.Pending
			},
			wantErr: true,
			errMsg:  "price must be at least 3000",
		},
		{
			name: "priceAtRequestが1000000を超える場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, "メッセージ", menteeID, 1000001, planID, contract_requestdm.Pending
			},
			wantErr: true,
			errMsg:  "priceAtRequest must be less than 1000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, message, menteeID, priceAtRequest, planID, isAccepted := tt.setupFunc(t)

			cr, err := contract_requestdm.NewContractRequest(id, message, menteeID, priceAtRequest, planID, isAccepted)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				return
			}

			require.NoError(t, err)
			if tt.assertions != nil {
				tt.assertions(t, cr, id, message, menteeID, priceAtRequest, planID, isAccepted)
			}
		})
	}
}

func TestContractRequest_NewContractRequestByVal(t *testing.T) {
	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted)
		assertions func(t *testing.T, cr *contract_requestdm.ContractRequest, id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted contract_requestdm.IsAccepted)
	}{
		{
			name: "DBから取得したデータでContractRequestを再構築できる",
			setupFunc: func(t *testing.T) (shared.UUID, string, shared.UUID, uint32, shared.UUID, contract_requestdm.IsAccepted) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				menteeID, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				planID, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440002")
				require.NoError(t, err)

				return id, "契約申請メッセージ", menteeID, 10000, planID, contract_requestdm.Accepted
			},
			assertions: func(t *testing.T, cr *contract_requestdm.ContractRequest, id shared.UUID, message string, menteeID shared.UUID, priceAtRequest uint32, planID shared.UUID, isAccepted contract_requestdm.IsAccepted) {
				assert.NotNil(t, cr)
				assert.Equal(t, id, cr.ID())
				assert.Equal(t, message, cr.Message())
				assert.Equal(t, menteeID, cr.MenteeID())
				assert.Equal(t, priceAtRequest, cr.PriceAtRequest())
				assert.Equal(t, planID, cr.PlanID())
				assert.Equal(t, isAccepted, cr.IsAccepted())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, message, menteeID, priceAtRequest, planID, isAccepted := tt.setupFunc(t)

			cr, err := contract_requestdm.NewContractRequestByVal(id, message, menteeID, priceAtRequest, planID, isAccepted)
			require.NoError(t, err)

			if tt.assertions != nil {
				tt.assertions(t, cr, id, message, menteeID, priceAtRequest, planID, isAccepted)
			}
		})
	}
}

func TestContractRequest_Getters(t *testing.T) {
	id := shared.NewUUID()
	message := "テストメッセージ"
	menteeID := shared.NewUUID()
	priceAtRequest := uint32(5000)
	planID := shared.NewUUID()
	isAccepted := contract_requestdm.Pending

	cr, err := contract_requestdm.NewContractRequest(id, message, menteeID, priceAtRequest, planID, isAccepted)
	require.NoError(t, err)

	t.Run("ID()はContractRequestのIDを返す", func(t *testing.T) {
		assert.Equal(t, id, cr.ID())
	})

	t.Run("Message()はMessageを返す", func(t *testing.T) {
		assert.Equal(t, message, cr.Message())
	})

	t.Run("MenteeID()はMenteeIDを返す", func(t *testing.T) {
		assert.Equal(t, menteeID, cr.MenteeID())
	})

	t.Run("PriceAtRequest()はPriceAtRequestを返す", func(t *testing.T) {
		assert.Equal(t, priceAtRequest, cr.PriceAtRequest())
	})

	t.Run("PlanID()はPlanIDを返す", func(t *testing.T) {
		assert.Equal(t, planID, cr.PlanID())
	})

	t.Run("IsAccepted()はIsAcceptedを返す", func(t *testing.T) {
		assert.Equal(t, isAccepted, cr.IsAccepted())
	})
}
