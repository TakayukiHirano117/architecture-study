# Issue #83 実装計画: メンター詳細画面の追加

## 調査結果

### 現状
- バックエンド：`MentorRecruitmentRepository` に `Store` のみ定義。`FindByID` が存在しない。
- アプリサービス：`create_mentor_recruitment_app_service.go` のみ。詳細取得サービスなし。
- コントローラ：`POST /mentor-recruitments` のみ。`GET /mentor-recruitments/:id` なし。
- フロントエンド：`frontend/src/app/` に基本テンプレートしかなく、メンター関連ページが一切存在しない。

### 参照した既存パターン
- `planapp.GetPlansAppService` + `PlanQueryService` + `GetPlansController` のパターンを踏襲する。
- DTO は `backend/src/core/infra/controllers/dto/plan_response.go` に倣う。

---

## 想定修正方針

DDD + オニオンアーキテクチャのレイヤー順（domain → app → infra）に沿って実装する。

---

## タスク分解

### Backend

#### 1. domain 層: リポジトリインターフェースに `FindByID` を追加
**ファイル:** `backend/src/core/domain/mentor_recruitmentdm/mentor_recruitment_repository.go`

```go
FindByID(ctx context.Context, id MentorRecruitmentID) (*MentorRecruitment, error)
```

#### 2. domain 層: クエリサービスインターフェースを新規作成
**ファイル:** `backend/src/core/app/mentorrecruitmentapp/mentor_recruitment_query_service.go`

プランの `PlanQueryService` に倣い、`MentorRecruitmentQueryService` インターフェースと
`GetMentorRecruitmentByID(ctx context.Context, id string) (*mentor_recruitmentdm.MentorRecruitment, error)` メソッドを定義する。

#### 3. app 層: 詳細取得アプリサービスを新規作成
**ファイル:** `backend/src/core/app/mentorrecruitmentapp/get_mentor_recruitment_app_service.go`

```go
type GetMentorRecruitmentAppService interface {
    Exec(ctx context.Context, id string) (*mentor_recruitmentdm.MentorRecruitment, error)
}
```

#### 4. infra 層: クエリサービス実装を新規作成
**ファイル:** `backend/src/core/infra/rdbimpl/mentor_recruitment_query_service_impl.go`

`mentor_recruitments` テーブルから指定 ID のレコードをSELECTし、タグJOINも含めてドメインオブジェクトに変換する。

```sql
SELECT mr.id, mr.user_id, mr.title, mr.description, mr.category_id,
       mr.consultation_type, mr.consultation_method,
       mr.budget_from, mr.budget_to, mr.application_period, mr.status,
       COALESCE(array_agg(t.id) FILTER (WHERE t.id IS NOT NULL), '{}') AS tag_ids,
       mr.created_at, mr.updated_at
FROM mentor_recruitments mr
LEFT JOIN mentor_recruitment_tags mrt ON mrt.mentor_recruitment_id = mr.id
LEFT JOIN tags t ON t.id = mrt.tag_id
WHERE mr.id = $1
GROUP BY mr.id
```

> **注意:** `mentor_recruitment_tags` テーブルの存在・スキーマをマイグレーションファイルで確認してから実装すること。
> タグ中間テーブルが存在しない場合は、タグなしのシンプルな SELECT から始める。

#### 5. infra 層: DTO を新規作成
**ファイル:** `backend/src/core/infra/controllers/dto/mentor_recruitment_response.go`

```go
type MentorRecruitmentResponse struct {
    ID                 string    `json:"id"`
    UserID             string    `json:"user_id"`
    Title              string    `json:"title"`
    Description        string    `json:"description"`
    CategoryID         string    `json:"category_id"`
    ConsultationType   string    `json:"consultation_type"`
    ConsultationMethod string    `json:"consultation_method"`
    BudgetFrom         uint32    `json:"budget_from"`
    BudgetTo           uint32    `json:"budget_to"`
    ApplicationPeriod  time.Time `json:"application_period"`
    Status             string    `json:"status"`
    CreatedAt          time.Time `json:"created_at"`
    UpdatedAt          time.Time `json:"updated_at"`
}
```

#### 6. infra 層: コントローラを新規作成
**ファイル:** `backend/src/core/infra/controllers/get_mentor_recruitment_controller.go`

`GET /mentor-recruitments/:id` を受け取り、アプリサービスを呼び出してJSONレスポンスを返す。

#### 7. infra 層: ルーティング登録
**ファイル:** `backend/src/core/infra/controllers/mentor_recruitment_controller.go`

```go
router.GET("/mentor-recruitments/:id", mc.getMentorRecruitmentController.Exec)
```

#### 8. mock 更新
**ファイル:** `backend/src/support/mock/domain/mentor_recruitmentdm/mentor_recruitment_repository_mock.go`

`go:generate` で mockgen を再生成するか、手動で `FindByID` のモックメソッドを追加する。

---

### Frontend

#### 9. メンター詳細ページを新規作成
**ファイル:** `frontend/src/app/mentor-recruitments/[id]/page.tsx`

- Next.js App Router の動的ルートセグメント `[id]` を使用。
- `params.id` を使いバックエンドの `GET /api/mentor-recruitments/:id` を `fetch` で呼び出す。
- タイトル・説明・予算・相談方法・カテゴリ・ステータスなどを表示する。

---

## リスク・注意点

| リスク | 対策 |
|--------|------|
| タグ中間テーブルのスキーマ不明 | マイグレーションファイルを先に確認し、存在しない場合はタグなしで実装してから拡張 |
| `MentorRecruitmentID` の型（UUID かどうか） | `mentor_recruitment_id.go` を確認済み（`shared.UUID` ベース）、文字列で受け取りパースする |
| フロントエンドの API URL 設定 | 既存の `fetch` パターンまたは環境変数（`NEXT_PUBLIC_API_URL`）を確認してから実装 |
| `application_period` の型 | ドメインモデルの `ApplicationPeriod` 型の実装（`time.Time` ラッパー）を確認すること |

## 実装順序

1. マイグレーションファイルでテーブルスキーマを確認
2. `ApplicationPeriod` 型の実装を確認
3. Backend: タスク 1 → 2 → 3 → 4 → 5 → 6 → 7 → 8
4. Frontend: タスク 9
5. `go test ./...` でテスト通過を確認
