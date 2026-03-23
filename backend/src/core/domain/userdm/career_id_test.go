package userdm_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestCareerId_NewCareerId(t *testing.T) {
	careerId := userdm.NewCareerID()

	assert.NotEmpty(t, careerId.String())

	_, err := uuid.Parse(careerId.String())
	assert.NoError(t, err, "生成されたIDは有効なUUIDであるべき")
}

func TestCareerId_NewCareerIdByVal(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効な値で作成できる",
			input:   "test-career-id",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			careerId, err := userdm.NewCareerIDByVal(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, careerId.String())
		})
	}
}

func TestCareerId_Equal(t *testing.T) {
	tests := []struct {
		name     string
		value1   string
		value2   string
		expected bool
	}{
		{
			name:     "同じ値は等しい",
			value1:   "test-career-id",
			value2:   "test-career-id",
			expected: true,
		},
		{
			name:     "異なる値は等しくない",
			value1:   "test-career-id-1",
			value2:   "test-career-id-2",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			careerId1, err := userdm.NewCareerIDByVal(tt.value1)
			require.NoError(t, err)

			careerId2, err := userdm.NewCareerIDByVal(tt.value2)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, careerId1.Equal(careerId2))
		})
	}
}
