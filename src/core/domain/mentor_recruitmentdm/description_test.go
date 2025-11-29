package mentor_recruitmentdm_test

import (
	"strings"
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDescription_NewDescription(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効な説明で作成できる",
			input:   "説明",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "2000文字以上はエラー",
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
			description, err := mentor_recruitmentdm.NewDescription(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.input, description.String())
		})
	}
}
