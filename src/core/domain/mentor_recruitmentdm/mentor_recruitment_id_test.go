package mentor_recruitmentdm_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
)

func TestMentorRecruitmentID_NewMentorRecruitmentID(t *testing.T) {
	id := mentor_recruitmentdm.NewMentorRecruitmentID()

	_, err := uuid.Parse(id.String())
	require.NoError(t, err)

	assert.NotEmpty(t, id.String())
}

func TestMentorRecruitmentID_NewMentorRecruitmentID_GeneratesUniqueIDs(t *testing.T) {
	id1 := mentor_recruitmentdm.NewMentorRecruitmentID()
	id2 := mentor_recruitmentdm.NewMentorRecruitmentID()

	assert.NotEqual(t, id1.String(), id2.String())
}

func TestMentorRecruitmentID_NewMentorRecruitmentIDByVal(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "有効なUUIDで作成できる",
			input:   "550e8400-e29b-41d4-a716-446655440000",
			wantErr: false,
		},
		{
			name:    "別の有効なUUIDで作成できる",
			input:   "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			wantErr: false,
		},
		{
			name:    "空文字はエラー",
			input:   "",
			wantErr: true,
		},
		{
			name:    "無効なUUID形式はエラー",
			input:   "invalid-uuid",
			wantErr: true,
		},
		{
			name:    "UUIDに似ているが無効な形式はエラー",
			input:   "550e8400-e29b-41d4-a716",
			wantErr: true,
		},
		{
			name:    "ただの文字列はエラー",
			input:   "hello-world",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := mentor_recruitmentdm.NewMentorRecruitmentIDByVal(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.input, id.String())
		})
	}
}

func TestMentorRecruitmentID_String(t *testing.T) {
	expectedUUID := "550e8400-e29b-41d4-a716-446655440000"
	id, err := mentor_recruitmentdm.NewMentorRecruitmentIDByVal(expectedUUID)
	require.NoError(t, err)

	assert.Equal(t, expectedUUID, id.String())
}

func TestMentorRecruitmentID_Equal(t *testing.T) {
	uuid1 := "550e8400-e29b-41d4-a716-446655440000"
	uuid2 := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"

	id1, err := mentor_recruitmentdm.NewMentorRecruitmentIDByVal(uuid1)
	require.NoError(t, err)

	id2, err := mentor_recruitmentdm.NewMentorRecruitmentIDByVal(uuid1)
	require.NoError(t, err)

	id3, err := mentor_recruitmentdm.NewMentorRecruitmentIDByVal(uuid2)
	require.NoError(t, err)

	assert.True(t, id1.Equal(id2))
	assert.False(t, id1.Equal(id3))
	assert.False(t, id2.Equal(id3))
}

func TestMentorRecruitmentID_Equal_WithNewID(t *testing.T) {
	// NewMentorRecruitmentIDで生成したIDとNewMentorRecruitmentIDByValで同じ値を作成した場合
	newID := mentor_recruitmentdm.NewMentorRecruitmentID()

	sameID, err := mentor_recruitmentdm.NewMentorRecruitmentIDByVal(newID.String())
	require.NoError(t, err)

	assert.True(t, newID.Equal(sameID))
}
