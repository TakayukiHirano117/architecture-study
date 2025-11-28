package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCareer_NewCareer(t *testing.T) {
	t.Run("有効なパラメータでCareerを作成できる", func(t *testing.T) {
		careerId := userdm.NewCareerID()

		careerDetail, err := userdm.NewCareerDetail("Web開発に従事")
		require.NoError(t, err)

		careerStartYear, err := userdm.NewCareerStartYear(2020)
		require.NoError(t, err)

		careerEndYear, err := userdm.NewCareerEndYear(2022)
		require.NoError(t, err)

		career, err := userdm.NewCareer(careerId, *careerDetail, *careerStartYear, *careerEndYear)

		require.NoError(t, err)
		assert.NotNil(t, career)
	})
}
