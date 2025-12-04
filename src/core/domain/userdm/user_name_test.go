package userdm_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestUserName_NewUserName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なユーザー名で作成できる",
			input:   "Test User",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "256文字以上はエラー",
			input:   strings.Repeat("a", 256),
			wantErr: true,
		},
		{
			name:    "境界値: 255文字は有効",
			input:   strings.Repeat("a", 255),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userName, err := userdm.NewUserName(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, string(*userName))
		})
	}
}
