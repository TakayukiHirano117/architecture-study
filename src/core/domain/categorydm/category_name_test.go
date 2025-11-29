package categorydm_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
)

func TestCategoryName_NewCategoryName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なカテゴリ名で作成できる",
			input:   "Test Category",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "256文字以上はエラー",
			input:   strings.Repeat("a", 256),
			wantErr: true,
		},
		{
			name:    "境界値: 255文字は有効",
			input:   strings.Repeat("a", 255),
			wantErr: false,
		},
		{
			name:    "日本語文字列で作成できる",
			input:   "プログラミング",
			wantErr: false,
		},
		{
			name:    "日本語255文字は有効",
			input:   strings.Repeat("あ", 255),
			wantErr: false,
		},
		{
			name:    "日本語256文字以上はエラー",
			input:   strings.Repeat("あ", 256),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			categoryName, err := categorydm.NewCategoryName(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, string(*categoryName))
		})
	}
}

func TestCategoryName_NewCategoryNameByVal(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なカテゴリ名で作成できる",
			input:   "Test Category",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "長い文字列でも作成できる（長さ制限なし）",
			input:   strings.Repeat("a", 300),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			categoryName, err := categorydm.NewCategoryNameByVal(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, categoryName.String())
		})
	}
}

func TestCategoryName_String(t *testing.T) {
	input := "Test Category"
	categoryName, err := categorydm.NewCategoryNameByVal(input)

	require.NoError(t, err)
	assert.Equal(t, input, categoryName.String())
}

func TestCategoryName_Equal(t *testing.T) {
	tests := []struct {
		name     string
		value1   string
		value2   string
		expected bool
	}{
		{
			name:     "同じ名前は等しい",
			value1:   "Category A",
			value2:   "Category A",
			expected: true,
		},
		{
			name:     "異なる名前は等しくない",
			value1:   "Category A",
			value2:   "Category B",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			categoryName1, err := categorydm.NewCategoryNameByVal(tt.value1)
			require.NoError(t, err)

			categoryName2, err := categorydm.NewCategoryNameByVal(tt.value2)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, categoryName1.Equal(categoryName2))
		})
	}
}
