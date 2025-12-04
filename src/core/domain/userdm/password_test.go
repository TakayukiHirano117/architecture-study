package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword_NewPassword(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なパスワードで作成できる",
			input:   "validPassword0",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "短すぎるパスワードはエラー",
			input:   "short1A",
			wantErr: true,
		},
		{
			name:    "英字がないパスワードはエラー",
			input:   "123456789012",
			wantErr: true,
		},
		{
			name:    "数字がないパスワードはエラー",
			input:   "validPasswordOnly",
			wantErr: true,
		},
		{
			name:    "境界値: 最小長のパスワードは有効",
			input:   "validPassw0rd",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := userdm.NewPassword(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			// パスワードはハッシュ化されるため、bcrypt.CompareHashAndPasswordで検証
			err = bcrypt.CompareHashAndPassword([]byte(password.String()), []byte(tt.input))
			assert.NoError(t, err, "ハッシュ化されたパスワードが元のパスワードと一致すること")
		})
	}
}
