package categorydm_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
)

func TestCategoryId_NewCategoryId(t *testing.T) {
	categoryId := categorydm.NewCategoryID()

	assert.NotEmpty(t, categoryId.String())

	_, err := uuid.Parse(categoryId.String())
	assert.NoError(t, err, "生成されたIDは有効なUUIDであるべき")
}

func TestCategoryId_NewCategoryIdByVal(t *testing.T) {
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
			categoryId, err := categorydm.NewCategoryIDByVal(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, categoryId.String())
		})
	}
}

func TestCategoryId_String(t *testing.T) {
	validUUID := uuid.New().String()
	categoryId, err := categorydm.NewCategoryIDByVal(validUUID)

	require.NoError(t, err)
	assert.Equal(t, validUUID, categoryId.String())
}

func TestCategoryId_Equal(t *testing.T) {
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
			categoryId1, err := categorydm.NewCategoryIDByVal(tt.value1)
			require.NoError(t, err)

			categoryId2, err := categorydm.NewCategoryIDByVal(tt.value2)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, categoryId1.Equal(categoryId2))
		})
	}
}
