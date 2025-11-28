package userdm_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestUserId_NewUserId(t *testing.T) {
	userId := userdm.NewUserID()

	assert.NotEmpty(t, userId.String())

	_, err := uuid.Parse(userId.String())
	assert.NoError(t, err, "生成されたIDは有効なUUIDであるべき")
}

func TestUserId_NewUserIdByVal(t *testing.T) {
	validUUID := uuid.New().String()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なUUIDで作成できる",
			input:   validUUID,
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "無効なUUID形式はエラー",
			input:   "invalid-uuid-string",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId, err := userdm.NewUserIDByVal(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, userId.String())
		})
	}
}

func TestUserId_Equal(t *testing.T) {
	validUUID := uuid.New().String()

	tests := []struct {
		name     string
		value1   string
		value2   string
		expected bool
	}{
		{
			name:     "同じUUIDは等しい",
			value1:   validUUID,
			value2:   validUUID,
			expected: true,
		},
		{
			name:     "異なるUUIDは等しくない",
			value1:   validUUID,
			value2:   uuid.New().String(),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId1, err := userdm.NewUserIDByVal(tt.value1)
			require.NoError(t, err)

			userId2, err := userdm.NewUserIDByVal(tt.value2)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, userId1.Equal(userId2))
		})
	}
}
