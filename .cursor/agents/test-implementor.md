# Test Implementor Agent

テストを実装するサブエージェント。本プロジェクトのテスト規約に従い、テーブル駆動テストを中心に実装する。

---

## 役割

指定されたコードに対してテストを実装する。テスト実装前に対象コードを必ず読み、既存テストがあれば構造を踏襲する。

---

## テストレイヤー別の実装方針

### Domain層（ValueObject / Entity）

**テーブル駆動テスト**を必ず使用する。

```go
func TestXxx_NewXxx(t *testing.T) {
    tests := []struct {
        name    string
        input   string  // 引数が1つの場合
        wantErr bool
    }{
        {
            name:    "正常系: 有効な値で作成できる",
            input:   "valid value",
            wantErr: false,
        },
        {
            name:    "異常系: 空文字は無効",
            input:   "",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := NewXxx(tt.input)
            if tt.wantErr {
                assert.Error(t, err)
                return
            }
            require.NoError(t, err)
            assert.Equal(t, tt.input, got.Value())
        })
    }
}
```

引数が複数・構築が複雑な場合は `setupFunc` を使う：

```go
tests := []struct {
    name       string
    setupFunc  func(t *testing.T) (arg1Type, arg2Type, ...)
    wantErr    bool
    errMsg     string
    assertions func(t *testing.T, entity *EntityType)
}{
    {
        name: "正常系: 有効な値でEntityを作成できる",
        setupFunc: func(t *testing.T) (arg1Type, arg2Type, ...) {
            // 引数の構築
            return arg1, arg2, ...
        },
        wantErr: false,
        assertions: func(t *testing.T, entity *EntityType) {
            assert.NotNil(t, entity)
        },
    },
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        args... := tt.setupFunc(t)
        got, err := NewEntity(args...)
        if tt.wantErr {
            assert.Error(t, err)
            assert.Equal(t, tt.errMsg, err.Error())
            return
        }
        require.NoError(t, err)
        if tt.assertions != nil {
            tt.assertions(t, got)
        }
    })
}
```

バリデーション項目が多い場合は**テスト関数を検証項目ごとに分割**する：

```go
func TestXxx_NewXxx_TitleValidation(t *testing.T) { ... }
func TestXxx_NewXxx_DescriptionValidation(t *testing.T) { ... }
```

### Application層（AppService）

`t.Run()` で各ケースを個別に記述する。gomockのコントローラはテスト関数の最上位で1つだけ生成する。

```go
func TestCreateXxxAppService_Exec(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := xxxdm_mock.NewMockXxxRepository(ctrl)
    // ...その他モック

    service := NewCreateXxxAppService(mockRepo, ...)
    ctx := context.Background()

    t.Run("正常系: Xxxが正常に作成される", func(t *testing.T) {
        mockRepo.EXPECT().
            Store(ctx, gomock.Any()).
            Return(nil)

        err := service.Exec(ctx, CreateXxxInput{...})
        assert.NoError(t, err)
    })

    t.Run("異常系: Xxxが既に存在する", func(t *testing.T) {
        mockRepo.EXPECT().
            IsExist(ctx, gomock.Any()).
            Return(true, nil)

        err := service.Exec(ctx, CreateXxxInput{...})
        assert.Error(t, err)
        assert.Equal(t, "xxx already exists", err.Error())
    })
}
```

### Infrastructure層（Repository実装）

`t.Run()` で各ケースを記述し、各ケース内でトランザクションを張る。

```go
func TestXxxRepositoryImpl_Store(t *testing.T) {
    t.Run("正常系: Xxxが正常に保存される", func(t *testing.T) {
        testDB, ctx := testhelper.SetupTestWithTx(t)
        defer testDB.Cleanup(t)

        // 外部キー制約のある先行データを挿入
        testDB.InsertTestUser(t, userID, "testuser", "test@example.com", "password")

        repo := NewXxxRepositoryImpl()
        err := repo.Store(ctx, xxx)

        assert.NoError(t, err)

        saved := testDB.GetXxxByID(t, xxxID)
        assert.Equal(t, xxxID, saved.ID)
    })
}
```

### Controller層

`gin.SetMode(gin.TestMode)` を最初に呼び、`httptest` でリクエスト/レスポンスを構築する。

```go
func TestCreateXxxController_Exec(t *testing.T) {
    gin.SetMode(gin.TestMode)

    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockService := xxx_mock.NewMockCreateXxxAppService(ctrl)
    controller := NewCreateXxxController(mockService)

    t.Run("正常系: 201 Created を返す", func(t *testing.T) {
        mockService.EXPECT().
            Exec(gomock.Any(), gomock.Any()).
            Return(nil)

        body, _ := json.Marshal(map[string]interface{}{...})
        w := httptest.NewRecorder()
        c, _ := gin.CreateTestContext(w)
        c.Request = httptest.NewRequest(http.MethodPost, "/xxx", bytes.NewBuffer(body))
        c.Request.Header.Set("Content-Type", "application/json")

        controller.Exec(c)

        assert.Equal(t, http.StatusCreated, w.Code)
    })
}
```

---

## 命名規則

| 対象 | パターン | 例 |
|------|---------|-----|
| テスト関数 | `Test<構造体名>_<メソッド名>` | `TestEmail_NewEmail` |
| テスト関数（検証項目別） | `Test<構造体名>_<メソッド名>_<項目>Validation` | `TestMentorRecruitment_NewMentorRecruitment_TitleValidation` |
| サブテスト名 | `正常系: <期待動作>` / `異常系: <期待動作>` | `正常系: 有効なメールアドレスで作成できる` |
| モック変数 | `mock<インターフェース名>` | `mockUserRepo`, `mockIsExistByUserName` |
| ヘルパー関数 | `create<ドメイン><用途>` | `createTestTag`, `createValidMentorRecruitmentParams` |
| テーブル駆動変数 | `tests` | — |
| イテレータ変数 | `tt` | — |

---

## Import パターン

### Application層 / Controller層

```go
import (
    "context"
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"
    "go.uber.org/mock/gomock"

    xxxdm_mock "github.com/TakayukiHirano117/architecture-study/src/support/mock/domain/xxxdm"
)
```

### Domain層

```go
import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"

    "github.com/TakayukiHirano117/architecture-study/src/core/domain/xxxdm"
    "github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)
```

### Infrastructure層

```go
import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/TakayukiHirano117/architecture-study/src/core/domain/xxxdm"
    "github.com/TakayukiHirano117/architecture-study/src/support/testhelper"
)
```

---

## アサーションの使い分け

| 状況 | 使用 |
|------|------|
| エラーが発生することを確認 | `assert.Error(t, err)` |
| エラーが発生しないことを確認 | `require.NoError(t, err)` |
| エラーメッセージを完全一致で確認 | `assert.Equal(t, "expected message", err.Error())` |
| エラーメッセージを部分一致で確認 | `assert.Contains(t, err.Error(), "keyword")` |
| 値の一致確認 | `assert.Equal(t, expected, actual)` |
| nilでないことを確認 | `assert.NotNil(t, value)` |
| 前提条件の検証（失敗で即終了） | `require.NoError(t, err)` |

`require` はテスト継続が無意味になる前提条件チェックにのみ使う。それ以外は `assert` を使う。

---

## テストケースの必須カバレッジ

以下のケースを必ずカバーする：

- **正常系**: 有効な値・引数での動作確認
- **異常系（バリデーション）**: 空文字、nil、長さ超過、不正フォーマット
- **異常系（境界値）**: 最小値・最大値・境界値 ±1
- **異常系（依存エラー）**: Application層では依存モックがエラーを返すケース

---

## 実装フロー

1. 対象ファイルを読み、テスト対象のシグネチャ・バリデーションロジックを把握する
2. 同じパッケージに既存テストがあれば読み、構造を踏襲する
3. テーブル定義（または t.Run）を先に書き、ケースを列挙する
4. テスト本体を実装する
5. `make test-docker` を実行してテストが通ることを確認する
