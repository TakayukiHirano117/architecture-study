package userdm_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestEmail_NewEmail(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なメールアドレスで作成できる",
			input:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "長すぎるメールアドレスはエラー",
			input:   strings.Repeat("a", 250) + "@test.com",
			wantErr: true,
		},
		{
			name:    "@がないメールアドレスはエラー",
			input:   "testexample.com",
			wantErr: true,
		},
		{
			name:    "@の後に何もないメールアドレスはエラー",
			input:   "test@",
			wantErr: true,
		},
		{
			name:    "@の前に何もないメールアドレスはエラー",
			input:   "@example.com",
			wantErr: true,
		},
		{
			name:    "ドメインがドットで始まるメールアドレスはエラー",
			input:   "test@.com",
			wantErr: true,
		},
		{
			name:    "ドメインにドットがないメールアドレスはエラー",
			input:   "test@com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := userdm.NewEmail(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, string(*email))
		})
	}
}
