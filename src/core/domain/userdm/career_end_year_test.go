package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCareerEndYear_NewCareerEndYear(t *testing.T) {
	tests := []struct {
		name    string
		input   uint16
		wantErr bool
	}{
		{
			name:    "有効な年で作成できる",
			input:   2020,
			wantErr: false,
		},
		{
			name:    "1970年未満はエラー",
			input:   1969,
			wantErr: true,
		},
		{
			name:    "境界値: 1970年は有効",
			input:   1970,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			careerEndYear, err := userdm.NewCareerEndYear(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, careerEndYear.Uint16())
		})
	}
}

func TestCareerEndYear_Equal(t *testing.T) {
	tests := []struct {
		name     string
		year1    uint16
		year2    uint16
		expected bool
	}{
		{
			name:     "同じ年は等しい",
			year1:    2020,
			year2:    2020,
			expected: true,
		},
		{
			name:     "異なる年は等しくない",
			year1:    2020,
			year2:    2021,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			careerEndYear1, err := userdm.NewCareerEndYear(tt.year1)
			require.NoError(t, err)

			careerEndYear2, err := userdm.NewCareerEndYear(tt.year2)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, careerEndYear1.Equal(*careerEndYear2))
		})
	}
}
