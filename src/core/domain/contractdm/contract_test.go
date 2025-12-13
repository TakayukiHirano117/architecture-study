package contractdm_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/contractdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

func TestContract_NewContract(t *testing.T) {
	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) (shared.UUID, shared.UUID, shared.UUID, string)
		wantErr    bool
		errMsg     string
		assertions func(t *testing.T, contract *contractdm.Contract, id, menteeID, planID shared.UUID, message string)
	}{
		{
			name: "有効な値でContractを作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, shared.UUID, shared.UUID, string) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, menteeID, planID, "プランに興味があります。よろしくお願いします。"
			},
			wantErr: false,
			assertions: func(t *testing.T, contract *contractdm.Contract, id, menteeID, planID shared.UUID, message string) {
				assert.NotNil(t, contract)
				assert.Equal(t, id, contract.ID())
				assert.Equal(t, menteeID, contract.MenteeID())
				assert.Equal(t, planID, contract.PlanID())
				assert.Equal(t, message, contract.Message())
			},
		},
		{
			name: "messageが500文字でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, shared.UUID, shared.UUID, string) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, menteeID, planID, strings.Repeat("あ", 500)
			},
			wantErr: false,
			assertions: func(t *testing.T, contract *contractdm.Contract, id, menteeID, planID shared.UUID, message string) {
				assert.NotNil(t, contract)
				assert.Equal(t, message, contract.Message())
			},
		},
		{
			name: "messageが1文字でも作成できる",
			setupFunc: func(t *testing.T) (shared.UUID, shared.UUID, shared.UUID, string) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, menteeID, planID, "a"
			},
			wantErr: false,
			assertions: func(t *testing.T, contract *contractdm.Contract, id, menteeID, planID shared.UUID, message string) {
				assert.NotNil(t, contract)
			},
		},
		{
			name: "messageが空の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, shared.UUID, shared.UUID, string) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, menteeID, planID, ""
			},
			wantErr: true,
			errMsg:  "message must not be empty",
		},
		{
			name: "messageが501文字以上の場合エラー",
			setupFunc: func(t *testing.T) (shared.UUID, shared.UUID, shared.UUID, string) {
				id := shared.NewUUID()
				menteeID := shared.NewUUID()
				planID := shared.NewUUID()
				return id, menteeID, planID, strings.Repeat("あ", 501)
			},
			wantErr: true,
			errMsg:  "message must be less than 500 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, menteeID, planID, message := tt.setupFunc(t)

			contract, err := contractdm.NewContract(id, menteeID, planID, message)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				return
			}

			require.NoError(t, err)
			if tt.assertions != nil {
				tt.assertions(t, contract, id, menteeID, planID, message)
			}
		})
	}
}

func TestContract_NewContractByVal(t *testing.T) {
	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) (shared.UUID, shared.UUID, shared.UUID, string)
		assertions func(t *testing.T, contract *contractdm.Contract, id, menteeID, planID shared.UUID, message string)
	}{
		{
			name: "DBから取得したデータでContractを再構築できる",
			setupFunc: func(t *testing.T) (shared.UUID, shared.UUID, shared.UUID, string) {
				id, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440000")
				require.NoError(t, err)

				menteeID, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440001")
				require.NoError(t, err)

				planID, err := shared.NewUUIDByVal("550e8400-e29b-41d4-a716-446655440002")
				require.NoError(t, err)

				return id, menteeID, planID, "契約メッセージ"
			},
			assertions: func(t *testing.T, contract *contractdm.Contract, id, menteeID, planID shared.UUID, message string) {
				assert.NotNil(t, contract)
				assert.Equal(t, id, contract.ID())
				assert.Equal(t, menteeID, contract.MenteeID())
				assert.Equal(t, planID, contract.PlanID())
				assert.Equal(t, message, contract.Message())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, menteeID, planID, message := tt.setupFunc(t)

			contract, err := contractdm.NewContractByVal(id, menteeID, planID, message)
			require.NoError(t, err)

			if tt.assertions != nil {
				tt.assertions(t, contract, id, menteeID, planID, message)
			}
		})
	}
}

func TestContract_Getters(t *testing.T) {
	id := shared.NewUUID()
	menteeID := shared.NewUUID()
	planID := shared.NewUUID()
	message := "テストメッセージ"

	contract, err := contractdm.NewContract(id, menteeID, planID, message)
	require.NoError(t, err)

	t.Run("ID()はContractのIDを返す", func(t *testing.T) {
		assert.Equal(t, id, contract.ID())
	})

	t.Run("MenteeID()はMenteeのIDを返す", func(t *testing.T) {
		assert.Equal(t, menteeID, contract.MenteeID())
	})

	t.Run("PlanID()はPlanのIDを返す", func(t *testing.T) {
		assert.Equal(t, planID, contract.PlanID())
	})

	t.Run("Message()はMessageを返す", func(t *testing.T) {
		assert.Equal(t, message, contract.Message())
	})
}
