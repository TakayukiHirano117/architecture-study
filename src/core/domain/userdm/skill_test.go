package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkill_NewSkill(t *testing.T) {
	t.Run("有効なパラメータでSkillを作成できる", func(t *testing.T) {
		skillId := userdm.NewSkillID()

		tagId := tagdm.NewTagID()
		tagName, err := tagdm.NewTagName("test-tag-name")
		require.NoError(t, err)

		tag, err := tagdm.NewTag(tagId, *tagName)
		require.NoError(t, err)

		evaluation, err := userdm.NewEvaluation(5)
		require.NoError(t, err)

		yearsOfExperience, err := userdm.NewYearsOfExperience(3)
		require.NoError(t, err)

		skill, err := userdm.NewSkill(skillId, tag, evaluation, yearsOfExperience)

		require.NoError(t, err)
		assert.NotNil(t, skill)
	})
}
