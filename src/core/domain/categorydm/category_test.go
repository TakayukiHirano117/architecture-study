package categorydm_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
)

func TestCategory_NewCategory(t *testing.T) {
	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) (categorydm.CategoryID, categorydm.CategoryName)
		wantErr    bool
		assertions func(t *testing.T, category *categorydm.Category, id categorydm.CategoryID, categoryName categorydm.CategoryName)
	}{
		{
			name: "有効なパラメータでCategoryを作成できる",
			setupFunc: func(t *testing.T) (categorydm.CategoryID, categorydm.CategoryName) {
				categoryId := categorydm.NewCategoryID()

				categoryName, err := categorydm.NewCategoryName("プログラミング")
				require.NoError(t, err)

				return categoryId, *categoryName
			},
			wantErr: false,
			assertions: func(t *testing.T, category *categorydm.Category, id categorydm.CategoryID, categoryName categorydm.CategoryName) {
				assert.NotNil(t, category)
				assert.Equal(t, id, category.ID())
				assert.Equal(t, categoryName, category.Name())
				assert.False(t, category.CreatedAt().IsZero())
				assert.False(t, category.UpdatedAt().IsZero())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beforeCreate := time.Now()
			categoryId, categoryName := tt.setupFunc(t)

			category, err := categorydm.NewCategory(categoryId, categoryName)
			afterCreate := time.Now()

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			if tt.assertions != nil {
				tt.assertions(t, category, categoryId, categoryName)
			}

			// CreatedAt と UpdatedAt が適切な範囲内であることを確認
			assert.True(t, !category.CreatedAt().Before(beforeCreate) && !category.CreatedAt().After(afterCreate),
				"CreatedAtは作成時刻の範囲内であるべき")
			assert.True(t, !category.UpdatedAt().Before(beforeCreate) && !category.UpdatedAt().After(afterCreate),
				"UpdatedAtは作成時刻の範囲内であるべき")
		})
	}
}

func TestCategory_NewCategoryByVal(t *testing.T) {
	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) (categorydm.CategoryID, categorydm.CategoryName, time.Time, time.Time)
		wantErr    bool
		assertions func(t *testing.T, category *categorydm.Category, id categorydm.CategoryID, categoryName categorydm.CategoryName)
	}{
		{
			name: "有効なパラメータでCategoryを作成できる",
			setupFunc: func(t *testing.T) (categorydm.CategoryID, categorydm.CategoryName, time.Time, time.Time) {
				categoryId := categorydm.NewCategoryID()

				categoryName, err := categorydm.NewCategoryNameByVal("デザイン")
				require.NoError(t, err)

				return categoryId, categoryName, time.Now(), time.Now()
			},
			wantErr: false,
			assertions: func(t *testing.T, category *categorydm.Category, id categorydm.CategoryID, categoryName categorydm.CategoryName) {
				assert.NotNil(t, category)
				assert.Equal(t, id, category.ID())
				assert.Equal(t, categoryName, category.Name())
				assert.False(t, category.CreatedAt().IsZero())
				assert.False(t, category.UpdatedAt().IsZero())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			categoryId, categoryName, createdAt, updatedAt := tt.setupFunc(t)

			category, err := categorydm.NewCategoryByVal(categoryId, categoryName, createdAt, updatedAt)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			if tt.assertions != nil {
				tt.assertions(t, category, categoryId, categoryName)
			}
		})
	}
}

func TestCategory_Getters(t *testing.T) {
	categoryId := categorydm.NewCategoryID()
	categoryName, err := categorydm.NewCategoryName("テストカテゴリ")
	require.NoError(t, err)

	category, err := categorydm.NewCategory(categoryId, *categoryName)
	require.NoError(t, err)

	t.Run("ID()はCategoryIDを返す", func(t *testing.T) {
		assert.Equal(t, categoryId, category.ID())
	})

	t.Run("Name()はCategoryNameを返す", func(t *testing.T) {
		assert.Equal(t, *categoryName, category.Name())
	})

	t.Run("CreatedAt()は作成日時を返す", func(t *testing.T) {
		assert.False(t, category.CreatedAt().IsZero())
	})

	t.Run("UpdatedAt()は更新日時を返す", func(t *testing.T) {
		assert.False(t, category.UpdatedAt().IsZero())
	})
}
