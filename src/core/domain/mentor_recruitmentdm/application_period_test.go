package mentor_recruitmentdm_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
)

func TestApplicationPeriod_NewApplicationPeriod(t *testing.T) {
	period := mentor_recruitmentdm.NewApplicationPeriod()

	// 今日から14日後の日付になっていること
	expected := time.Now().AddDate(0, 0, 14)
	expectedDate := time.Date(expected.Year(), expected.Month(), expected.Day(), 0, 0, 0, 0, expected.Location())

	assert.Equal(t, expectedDate, period.Time())
}

func TestApplicationPeriod_NewApplicationPeriodByVal(t *testing.T) {
	tests := []struct {
		name    string
		input   time.Time
		wantErr bool
	}{
		{
			name:    "有効な日付で作成できる",
			input:   time.Date(2025, 12, 2, 15, 30, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name:    "ゼロ値はエラー",
			input:   time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			period, err := mentor_recruitmentdm.NewApplicationPeriodByVal(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			// 時刻が00:00:00になっていること
			assert.Equal(t, 0, period.Time().Hour())
			assert.Equal(t, 0, period.Time().Minute())
			assert.Equal(t, 0, period.Time().Second())
			// 日付は保持されていること
			assert.Equal(t, tt.input.Year(), period.Time().Year())
			assert.Equal(t, tt.input.Month(), period.Time().Month())
			assert.Equal(t, tt.input.Day(), period.Time().Day())
		})
	}
}

func TestApplicationPeriod_String(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "12月02日形式で取得できる",
			input:    time.Date(2025, 12, 2, 0, 0, 0, 0, time.Local),
			expected: "12月02日",
		},
		{
			name:     "1月15日形式で取得できる",
			input:    time.Date(2025, 1, 15, 0, 0, 0, 0, time.Local),
			expected: "1月15日",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			period, err := mentor_recruitmentdm.NewApplicationPeriodByVal(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, period.String())
		})
	}
}

func TestApplicationPeriod_Equal(t *testing.T) {
	tests := []struct {
		name     string
		input1   time.Time
		input2   time.Time
		expected bool
	}{
		{
			name:     "同じ日付はtrue",
			input1:   time.Date(2025, 12, 2, 0, 0, 0, 0, time.Local),
			input2:   time.Date(2025, 12, 2, 0, 0, 0, 0, time.Local),
			expected: true,
		},
		{
			name:     "時刻が違っても同じ日付ならtrue",
			input1:   time.Date(2025, 12, 2, 0, 0, 0, 0, time.Local),
			input2:   time.Date(2025, 12, 2, 15, 30, 0, 0, time.Local),
			expected: true,
		},
		{
			name:     "違う日付はfalse",
			input1:   time.Date(2025, 12, 2, 0, 0, 0, 0, time.Local),
			input2:   time.Date(2025, 12, 3, 0, 0, 0, 0, time.Local),
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			period1, err := mentor_recruitmentdm.NewApplicationPeriodByVal(tt.input1)
			require.NoError(t, err)
			period2, err := mentor_recruitmentdm.NewApplicationPeriodByVal(tt.input2)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, period1.Equal(period2))
		})
	}
}
