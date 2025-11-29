package mentor_recruitmentdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBudget_NewBudget(t *testing.T) {
	tests := []struct {
		name    string
		min     uint32
		max     uint32
		wantErr bool
	}{
		{
			name:    "有効な予算範囲で作成できる",
			min:     3000,
			max:     10000,
			wantErr: false,
		},
		{
			name:    "最小値と最大値が同じでも作成できる",
			min:     5000,
			max:     5000,
			wantErr: false,
		},
		{
			name:    "最小値が1000未満はエラー",
			min:     999,
			max:     10000,
			wantErr: true,
		},
		{
			name:    "最大値が1000000を超える場合はエラー",
			min:     3000,
			max:     1000001,
			wantErr: true,
		},
		{
			name:    "最小値が最大値を超える場合はエラー",
			min:     10000,
			max:     3000,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			budget, err := mentor_recruitmentdm.NewBudget(tt.min, tt.max)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.min, budget.Min())
			assert.Equal(t, tt.max, budget.Max())
		})
	}
}

func TestBudget_Equal(t *testing.T) {
	budget1, _ := mentor_recruitmentdm.NewBudget(3000, 10000)
	budget2, _ := mentor_recruitmentdm.NewBudget(3000, 10000)
	budget3, _ := mentor_recruitmentdm.NewBudget(5000, 10000)

	assert.True(t, budget1.Equal(budget2))
	assert.False(t, budget1.Equal(budget3))
}
