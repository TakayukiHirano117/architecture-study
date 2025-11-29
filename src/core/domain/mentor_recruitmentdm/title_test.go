package mentor_recruitmentdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTitle_NewTitle(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なタイトルで作成できる",
			input:   "タイトル",
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
			title, err := mentor_recruitmentdm.NewTitle(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.input, title.String())
		})
	}
}
