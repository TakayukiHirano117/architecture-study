package mentor_recruitmentdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConsultationType_NewConsultationType(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "単発で作成できる",
			input:   "単発",
			wantErr: false,
		},
		{
			name:    "継続で作成できる",
			input:   "継続",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "無効な値はエラー",
			input:   "定期",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			consultationType, err := mentor_recruitmentdm.NewConsultationType(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			str, err := consultationType.String()
			require.NoError(t, err)
			assert.Equal(t, tt.input, str)
		})
	}
}

func TestConsultationType_NewConsultationTypeByVal(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効な値で作成できる",
			input:   "単発",
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
			consultationType, err := mentor_recruitmentdm.NewConsultationTypeByVal(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			str, err := consultationType.String()
			require.NoError(t, err)
			assert.Equal(t, tt.input, str)
		})
	}
}

func TestConsultationType_String(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "単発の文字列を取得できる",
			input:    "単発",
			expected: "単発",
		},
		{
			name:     "継続の文字列を取得できる",
			input:    "継続",
			expected: "継続",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			consultationType, _ := mentor_recruitmentdm.NewConsultationType(tt.input)
			str, err := consultationType.String()
			require.NoError(t, err)
			assert.Equal(t, tt.expected, str)
		})
	}
}

func TestConsultationType_String_InvalidType(t *testing.T) {
	invalidType, err := mentor_recruitmentdm.NewConsultationTypeByVal("無効")
	require.NoError(t, err)
	_, err = invalidType.String()
	assert.Error(t, err)
}

func TestConsultationType_Equal(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   string
		expected bool
	}{
		{
			name:     "同じ値はtrue",
			input1:   "単発",
			input2:   "単発",
			expected: true,
		},
		{
			name:     "違う値はfalse",
			input1:   "単発",
			input2:   "継続",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type1, _ := mentor_recruitmentdm.NewConsultationType(tt.input1)
			type2, _ := mentor_recruitmentdm.NewConsultationType(tt.input2)
			assert.Equal(t, tt.expected, type1.Equal(type2))
		})
	}
}
