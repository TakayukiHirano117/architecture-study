package rdbimpl

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/support/testhelper"
)

func TestPlanRepositoryImpl_Store(t *testing.T) {
	t.Run("正常系: タグなしでプランが正常に保存される", func(t *testing.T) {
		// テスト用のDB接続とトランザクションをセットアップ
		testDB, ctx := testhelper.SetupTestWithTx(t)
		defer testDB.Cleanup(t)

		// テストデータの準備
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		// 依存データの挿入
		testDB.InsertTestUser(t, userID, "testuser", "test@example.com", "password123456")
		testDB.InsertTestCategory(t, categoryID, "テストカテゴリ")

		// プランの作成
		planID := shared.NewUUID()
		mentorID, _ := userdm.NewUserIDByVal(userID)
		catID, _ := categorydm.NewCategoryIDByVal(categoryID)
		status, _ := plandm.NewStatus("公開")
		consultationType, _ := plandm.NewConsultationType("単発")

		plan, err := plandm.NewPlan(
			planID,
			mentorID,
			"テストプラン",
			catID,
			[]shared.UUID{},
			"テストプランの説明です。",
			status,
			&consultationType,
			10000,
		)
		assert.NoError(t, err)

		// リポジトリの実行
		repo := NewPlanRepositoryImpl()
		err = repo.Store(ctx, plan)

		// 検証
		assert.NoError(t, err)

		// DBからデータを取得して検証
		savedPlan := testDB.GetPlanByID(t, planID.String())
		assert.Equal(t, planID.String(), savedPlan["id"])
		assert.Equal(t, userID, savedPlan["mentor_id"])
		assert.Equal(t, "テストプラン", savedPlan["title"])
		assert.Equal(t, categoryID, savedPlan["category_id"])
		assert.Equal(t, "テストプランの説明です。", savedPlan["description"])
		assert.Equal(t, "公開", savedPlan["status"])
		assert.Equal(t, "単発", savedPlan["consultation_type"])
		assert.Equal(t, int64(10000), savedPlan["price"])
	})

	t.Run("正常系: タグ付きでプランが正常に保存される", func(t *testing.T) {
		// テスト用のDB接続とトランザクションをセットアップ
		testDB, ctx := testhelper.SetupTestWithTx(t)
		defer testDB.Cleanup(t)

		// テストデータの準備
		userID := uuid.New().String()
		categoryID := uuid.New().String()
		tagID1 := uuid.New().String()
		tagID2 := uuid.New().String()

		// 依存データの挿入（タグ名はユニークにするためUUIDを含める）
		testDB.InsertTestUser(t, userID, "testuser", "test@example.com", "password123456")
		testDB.InsertTestCategory(t, categoryID, "テストカテゴリ")
		testDB.InsertTestTag(t, tagID1, "TestTag_"+tagID1[:8])
		testDB.InsertTestTag(t, tagID2, "TestTag_"+tagID2[:8])

		// プランの作成
		planID := shared.NewUUID()
		mentorID, _ := userdm.NewUserIDByVal(userID)
		catID, _ := categorydm.NewCategoryIDByVal(categoryID)
		status, _ := plandm.NewStatus("公開")
		consultationType, _ := plandm.NewConsultationType("継続")

		tag1, _ := shared.NewUUIDByVal(tagID1)
		tag2, _ := shared.NewUUIDByVal(tagID2)
		tagIDs := []shared.UUID{tag1, tag2}

		plan, err := plandm.NewPlan(
			planID,
			mentorID,
			"Goメンタリングプラン",
			catID,
			tagIDs,
			"Go言語の基礎から応用までサポートします。",
			status,
			&consultationType,
			50000,
		)
		assert.NoError(t, err)

		// リポジトリの実行
		repo := NewPlanRepositoryImpl()
		err = repo.Store(ctx, plan)

		// 検証
		assert.NoError(t, err)

		// DBからデータを取得して検証
		savedPlan := testDB.GetPlanByID(t, planID.String())
		assert.Equal(t, planID.String(), savedPlan["id"])
		assert.Equal(t, "Goメンタリングプラン", savedPlan["title"])
		assert.Equal(t, "継続", savedPlan["consultation_type"])
		assert.Equal(t, int64(50000), savedPlan["price"])

		// タグの検証
		savedTagIDs := testDB.GetPlanTagsByPlanID(t, planID.String())
		assert.Len(t, savedTagIDs, 2)
		assert.Contains(t, savedTagIDs, tagID1)
		assert.Contains(t, savedTagIDs, tagID2)
	})

	t.Run("異常系: 存在しないユーザーIDでエラー", func(t *testing.T) {
		// テスト用のDB接続とトランザクションをセットアップ
		testDB, ctx := testhelper.SetupTestWithTx(t)
		defer testDB.Cleanup(t)

		// テストデータの準備（ユーザーは挿入しない）
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		// カテゴリのみ挿入
		testDB.InsertTestCategory(t, categoryID, "テストカテゴリ")

		// プランの作成
		planID := shared.NewUUID()
		mentorID, _ := userdm.NewUserIDByVal(userID)
		catID, _ := categorydm.NewCategoryIDByVal(categoryID)
		status, _ := plandm.NewStatus("公開")
		consultationType, _ := plandm.NewConsultationType("単発")

		plan, err := plandm.NewPlan(
			planID,
			mentorID,
			"テストプラン",
			catID,
			[]shared.UUID{},
			"テストプランの説明です。",
			status,
			&consultationType,
			10000,
		)
		assert.NoError(t, err)

		// リポジトリの実行
		repo := NewPlanRepositoryImpl()
		err = repo.Store(ctx, plan)

		// 外部キー制約違反でエラーになることを検証
		assert.Error(t, err)
	})

	t.Run("異常系: 存在しないカテゴリIDでエラー", func(t *testing.T) {
		// テスト用のDB接続とトランザクションをセットアップ
		testDB, ctx := testhelper.SetupTestWithTx(t)
		defer testDB.Cleanup(t)

		// テストデータの準備（カテゴリは挿入しない）
		userID := uuid.New().String()
		categoryID := uuid.New().String()

		// ユーザーのみ挿入
		testDB.InsertTestUser(t, userID, "testuser", "test@example.com", "password123456")

		// プランの作成
		planID := shared.NewUUID()
		mentorID, _ := userdm.NewUserIDByVal(userID)
		catID, _ := categorydm.NewCategoryIDByVal(categoryID)
		status, _ := plandm.NewStatus("公開")
		consultationType, _ := plandm.NewConsultationType("単発")

		plan, err := plandm.NewPlan(
			planID,
			mentorID,
			"テストプラン",
			catID,
			[]shared.UUID{},
			"テストプランの説明です。",
			status,
			&consultationType,
			10000,
		)
		assert.NoError(t, err)

		// リポジトリの実行
		repo := NewPlanRepositoryImpl()
		err = repo.Store(ctx, plan)

		// 外部キー制約違反でエラーになることを検証
		assert.Error(t, err)
	})

	t.Run("異常系: 存在しないタグIDでエラー", func(t *testing.T) {
		// テスト用のDB接続とトランザクションをセットアップ
		testDB, ctx := testhelper.SetupTestWithTx(t)
		defer testDB.Cleanup(t)

		// テストデータの準備
		userID := uuid.New().String()
		categoryID := uuid.New().String()
		tagID := uuid.New().String() // 存在しないタグID

		// ユーザーとカテゴリを挿入（タグは挿入しない）
		testDB.InsertTestUser(t, userID, "testuser", "test@example.com", "password123456")
		testDB.InsertTestCategory(t, categoryID, "テストカテゴリ")

		// プランの作成
		planID := shared.NewUUID()
		mentorID, _ := userdm.NewUserIDByVal(userID)
		catID, _ := categorydm.NewCategoryIDByVal(categoryID)
		status, _ := plandm.NewStatus("公開")
		consultationType, _ := plandm.NewConsultationType("単発")

		tag, _ := shared.NewUUIDByVal(tagID)
		tagIDs := []shared.UUID{tag}

		plan, err := plandm.NewPlan(
			planID,
			mentorID,
			"テストプラン",
			catID,
			tagIDs,
			"テストプランの説明です。",
			status,
			&consultationType,
			10000,
		)
		assert.NoError(t, err)

		// リポジトリの実行
		repo := NewPlanRepositoryImpl()
		err = repo.Store(ctx, plan)

		// 外部キー制約違反でエラーになることを検証
		assert.Error(t, err)
	})
}
