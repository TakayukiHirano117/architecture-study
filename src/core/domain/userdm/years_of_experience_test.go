package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestYearsOfExperience_NewYearsOfExperience(t *testing.T) {
	tests := []struct {
		name    string
		input   uint8
		wantErr bool
	}{
		{
			name:    "有効な経験年数で作成できる",
			input:   3,
			wantErr: false,
		},
		{
			name:    "100を超える経験年数はエラー",
			input:   101,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yearsOfExperience, err := userdm.NewYearsOfExperience(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, yearsOfExperience.Uint8())
		})
	}
}
