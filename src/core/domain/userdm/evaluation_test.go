package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvaluation_NewEvaluation(t *testing.T) {
	tests := []struct {
		name    string
		input   uint8
		wantErr bool
	}{
		{
			name:    "有効な評価値で作成できる",
			input:   5,
			wantErr: false,
		},
		{
			name:    "6以上はエラー",
			input:   6,
			wantErr: true,
		},
		{
			name:    "境界値: 0は有効",
			input:   0,
			wantErr: false,
		},
		{
			name:    "境界値: 5は有効",
			input:   5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluation, err := userdm.NewEvaluation(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, evaluation.Uint8())
		})
	}
}
