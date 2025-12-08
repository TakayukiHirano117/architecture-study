package shared_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

func TestUUID_NewUUID(t *testing.T) {
	uuid1 := shared.NewUUID()

	assert.NotEmpty(t, uuid1.String())

	_, err := uuid.Parse(uuid1.String())
	assert.NoError(t, err, "生成されたIDは有効なUUIDであるべき")

	uuid2 := shared.NewUUID()
	assert.NotEqual(t, uuid1, uuid2, "生成されるUUIDは毎回異なるべき")
}

func TestUUID_NewUUIDByVal(t *testing.T) {
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
		{
			name:    "不完全なUUID形式はエラー",
			input:   "550e8400-e29b-41d4",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uuid, err := shared.NewUUIDByVal(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, uuid.String())
		})
	}
}

func TestUUID_String(t *testing.T) {
	validUUID := uuid.New().String()
	u, err := shared.NewUUIDByVal(validUUID)

	require.NoError(t, err)
	assert.Equal(t, validUUID, u.String())
}

func TestUUID_Equal(t *testing.T) {
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
			uuid1, err := shared.NewUUIDByVal(tt.value1)
			require.NoError(t, err)

			uuid2, err := shared.NewUUIDByVal(tt.value2)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, uuid1.Equal(uuid2))
		})
	}
}
