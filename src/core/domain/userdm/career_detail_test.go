package userdm_test

import (
	"strings"
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCareerDetail_NewCareerDetail(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効な経歴詳細で作成できる",
			input:   "Webアプリケーション開発に従事しました。",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "2001文字以上はエラー",
			input:   strings.Repeat("a", 2001),
			wantErr: true,
		},
		{
			name:    "境界値: 2000文字は有効",
			input:   strings.Repeat("a", 2000),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			careerDetail, err := userdm.NewCareerDetail(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, string(*careerDetail))
		})
	}
}
