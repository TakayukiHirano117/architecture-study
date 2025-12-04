package mentor_recruitmentdm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
)

func TestConsultationMethod_NewConsultationMethod(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "チャットで作成できる",
			input:   "チャット",
			wantErr: false,
		},
		{
			name:    "ビデオ通話で作成できる",
			input:   "ビデオ通話",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "無効な値はエラー",
			input:   "電話",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method, err := mentor_recruitmentdm.NewConsultationMethod(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			str, err := method.String()
			require.NoError(t, err)
			assert.Equal(t, tt.input, str)
		})
	}
}

func TestConsultationMethod_NewConsultationMethodByVal(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効な値で作成できる",
			input:   "チャット",
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
			method, err := mentor_recruitmentdm.NewConsultationMethodByVal(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			str, err := method.String()
			require.NoError(t, err)
			assert.Equal(t, tt.input, str)
		})
	}
}

func TestConsultationMethod_String(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "チャットの文字列を取得できる",
			input:    "チャット",
			expected: "チャット",
		},
		{
			name:     "ビデオ通話の文字列を取得できる",
			input:    "ビデオ通話",
			expected: "ビデオ通話",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method, err := mentor_recruitmentdm.NewConsultationMethod(tt.input)
			require.NoError(t, err)
			str, err := method.String()
			require.NoError(t, err)
			assert.Equal(t, tt.expected, str)
		})
	}
}

func TestConsultationMethod_String_InvalidMethod(t *testing.T) {
	invalidMethod, err := mentor_recruitmentdm.NewConsultationMethodByVal("無効")
	require.NoError(t, err)
	_, err = invalidMethod.String()
	assert.Error(t, err)
}

func TestConsultationMethod_Equal(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   string
		expected bool
	}{
		{
			name:     "同じ値はtrue",
			input1:   "チャット",
			input2:   "チャット",
			expected: true,
		},
		{
			name:     "違う値はfalse",
			input1:   "チャット",
			input2:   "ビデオ通話",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method1, err := mentor_recruitmentdm.NewConsultationMethod(tt.input1)
			require.NoError(t, err)
			method2, err := mentor_recruitmentdm.NewConsultationMethod(tt.input2)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, method1.Equal(method2))
		})
	}
}
