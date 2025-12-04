package mentor_recruitmentdm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
)

func TestStatus_NewStatus(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "公開ステータスで作成できる",
			input:   "公開",
			wantErr: false,
		},
		{
			name:    "中止ステータスで作成できる",
			input:   "中止",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "無効な値はエラー",
			input:   "下書き",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, err := mentor_recruitmentdm.NewStatus(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			str, err := status.String()
			require.NoError(t, err)
			assert.Equal(t, tt.input, str)
		})
	}
}

func TestStatus_NewStatusByVal(t *testing.T) {
	status, err := mentor_recruitmentdm.NewStatusByVal("公開")
	require.NoError(t, err)
	str, err := status.String()
	require.NoError(t, err)
	assert.Equal(t, "公開", str)
}

func TestStatus_String(t *testing.T) {
	published, err := mentor_recruitmentdm.NewStatus("公開")
	require.NoError(t, err)

	cancelled, err := mentor_recruitmentdm.NewStatus("中止")
	require.NoError(t, err)

	publishedStr, err := published.String()
	require.NoError(t, err)
	assert.Equal(t, "公開", publishedStr)

	cancelledStr, err := cancelled.String()
	require.NoError(t, err)
	assert.Equal(t, "中止", cancelledStr)
}

func TestStatus_String_InvalidStatus(t *testing.T) {
	invalidStatus, err := mentor_recruitmentdm.NewStatusByVal("無効")
	require.NoError(t, err)
	_, err = invalidStatus.String()
	assert.Error(t, err)
}

func TestStatus_Equal(t *testing.T) {
	status1, err := mentor_recruitmentdm.NewStatus("公開")
	require.NoError(t, err)
	status2, err := mentor_recruitmentdm.NewStatus("公開")
	require.NoError(t, err)
	status3, err := mentor_recruitmentdm.NewStatus("中止")
	require.NoError(t, err)

	assert.True(t, status1.Equal(status2))
	assert.False(t, status1.Equal(status3))
	assert.False(t, status2.Equal(status3))
}
