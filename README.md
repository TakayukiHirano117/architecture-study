# ユースケース　
- メールパスログイン→こちらはwant
- ユーザ新規作成
    - 必須項目
        - 名前: 255文字以内(ユニーク)
        - メールアドレス: 255文字以内
        - パスワード: 12文字以上,英数字それぞれ最低1文字以上
        - スキル(複数)
            - タグ名(選択式、記入どちらもOK)
            - 評価: 1~5
            - 1つ以上
            - 年数: ０以上のint型(5年まで)
    - 任意項目
        - 自己紹介: 2000字以内
        - 経歴(複数)
            - 詳細: 1000字以内
            - 西暦from: 1970年以上のint型
            - 西暦to: 1970年以上のint型、西暦fromよりも大きい数字
- ユーザ更新
    - ユーザ新規作成と同様
- メンター募集作成
    - 必須項目
        - タイトル: 255文字以内
        - カテゴリ(1つ)
            - プログラミング
            - マーケティング
            - デザイン
            - ライティング
            - 動画・映像
            - ビジネス
            - 語学
            - ライフスタイル
        - 相談形式
            - 単発
            - 継続
        - 相談方式
            - チャット
            - ビデオ通話
        - ディスクリプション(2000文字)
        - 予算(from円~to円)
            - 1000円以上の数値
        - 応募期間
            - 1日単位
            - 最大14日後まで可能
        - ステータス(公開or中止)
    - 任意項目
        - タグ(複数OK、選択式、記入どちらもOK)
- メンター募集一覧
    - 表示するもの
        - タイトル
        - タグ一覧
        - カテゴリ
        - 予算
        - 提案数
        - 応募期間
            - 単発or継続
        - 提案したユーザ名
    - ページネーション20件
- メンター募集に提案する
    - 提案内容: 2000文字以内
    - 提案できるのはスキルを5個以上持つユーザのみ
- プラン作成
    - 以下必須項目
        - タイトル: 255文字以内
        - カテゴリ(1つ)
            - 上記のカテゴリ参照
        - タグ(5個まで)
            - 選択式
            - 新規記入
        - 内容
            - 2000文字以内
        - ステータス
            - 公開or中止
        - 相談形式
            - 単発
            - 継続
        - 金額
    - 任意項目
        - 相談方式
            - チャットorビデオ電話
- プラン契約リクエストを送る
    - メッセージ(最大500文字)
- プランへの契約を承認する
    - メッセージ(最大500文字)
- プラン一覧
    - 表示するもの
        - タイトル
        - カテゴリ
        - タグ
        - 内容
        - ステータス
        - 料金体系
        - 相談方式
        - プランを作成したユーザ名
    - ステータスが「中止」となっていないものだけ表示する
    - ページネーション20件
    

### 別途仕様で考慮するべきこと

- 月額契約は2ヶ月目以降は契約時の金額を請求する
    - 例えば1ヶ月目に1万円で契約したが2ヶ月目はプランが2万円の金額に変わっていた場合、引き続き1万円で契約は続行する
    - 2万円に変わってから契約した人は2万円の月額契約となる

## テストを書くもの
- Domain
  - Entity
  - vo
- Infra
  - QueryService
  - DomainService
  - Repository
- Usecase
ControllerはFW依存になるので作らない。

## このアプリの設計・アーキテクチャについて
DDD, オニオンアーキテクチャを使用しています。

DomainServiceはinterfaceと実装をいずれもDomain層に、

ControllerはInfra層に配置しています。

テストはテーブル駆動テスト方式を採用しています。

使用しているパッケージはgo.modを見てください。

CI/CDにGitHub Actionsを使用しています。

開発環境はDockerを用いています。

トランザクションの処理は少し特殊で、

HTTPメソッドにより自動的にDB接続かトランザクションを張ったDB接続を生成し、

自動でcommit, rollbackまで行ってくれる様にしています。

## プロジェクト詳細

### 使用技術

| カテゴリ | 技術 | バージョン/備考 |
|---------|------|----------------|
| 言語 | Go | 1.24+ |
| Webフレームワーク | Gin | gin-gonic/gin |
| データベース | PostgreSQL | 17.1-alpine (Docker) |
| DB接続 | sqlx | jmoiron/sqlx |
| パスワードハッシュ | bcrypt | golang.org/x/crypto |
| UUID | google/uuid | - |
| エラーハンドリング | cockroachdb/errors | スタックトレース付きエラー |
| テスト | testify, gomock | go.uber.org/mock |
| Lint | golangci-lint | - |
| ホットリロード | Air | air-verse/air |
| マイグレーション | golang-migrate | - |
| コンテナ | Docker, Docker Compose | - |
| CI/CD | GitHub Actions | format → lint → test |

### ディレクトリ構成

```
.
├── .docker/                    # Docker関連ファイル
│   ├── Dockerfile
│   └── compose.yml
├── .github/
│   └── workflows/
│       └── ci.yml              # GitHub Actions CI設定
├── config/
│   └── db.go                   # DB設定・環境変数読み込み
├── src/
│   ├── core/                   # コアロジック（オニオンアーキテクチャ）
│   │   ├── app/                # アプリケーション層（ユースケース）
│   │   │   ├── contractapp/
│   │   │   ├── mentorrecruitmentapp/
│   │   │   ├── planapp/
│   │   │   ├── query_service/
│   │   │   └── userapp/
│   │   ├── cmd/
│   │   │   └── main.go         # エントリーポイント
│   │   ├── domain/             # ドメイン層（ビジネスロジック）
│   │   │   ├── categorydm/     # カテゴリ集約
│   │   │   ├── contractdm/     # 契約集約
│   │   │   ├── mentor_recruitmentdm/  # メンター募集集約
│   │   │   ├── plandm/         # プラン集約
│   │   │   ├── shared/         # 共通値オブジェクト（UUID等）
│   │   │   ├── tagdm/          # タグ集約
│   │   │   └── userdm/         # ユーザー集約
│   │   └── infra/              # インフラ層（外部依存）
│   │       ├── controllers/    # HTTPコントローラー
│   │       ├── middlewares/    # Ginミドルウェア
│   │       ├── models/         # DBモデル（DTO）
│   │       ├── rdb/            # RDBハンドラーインターフェース
│   │       └── rdbimpl/        # リポジトリ実装
│   ├── db/
│   │   ├── init.sql            # 初期スキーマ
│   │   ├── migrations/         # マイグレーションファイル
│   │   └── rdb_handler.go      # RDBHandlerインターフェース
│   └── support/
│       ├── customerr/          # カスタムエラー定義
│       ├── mock/               # gomock生成ファイル
│       └── testhelper/         # テストヘルパー
├── go.mod
├── go.sum
├── Makefile                    # 開発用コマンド
└── README.md
```

### アーキテクチャ詳細

#### レイヤー構成（オニオンアーキテクチャ）

```
┌─────────────────────────────────────────────────────────┐
│                    Infra層（外側）                        │
│  Controllers, Middlewares, RDBImpl, Models              │
├─────────────────────────────────────────────────────────┤
│                 Application層（中間）                     │
│  AppService（ユースケース）, QueryService                 │
├─────────────────────────────────────────────────────────┤
│                   Domain層（内側）                        │
│  Entity, ValueObject, Repository(interface),            │
│  DomainService(interface + impl)                        │
└─────────────────────────────────────────────────────────┘
```

#### 依存の方向

- 外側 → 内側 への依存のみ許可
- Domain層は他の層に依存しない
- Repository/DomainServiceはDomain層でインターフェース定義、Infra層で実装

#### ドメイン層の構成

各集約（Aggregate）は以下の構成を持つ：

| ファイル | 役割 |
|---------|------|
| `{entity}.go` | エンティティ（集約ルート） |
| `{value_object}.go` | 値オブジェクト |
| `{entity}_repository.go` | リポジトリインターフェース |
| `is_exist_by_{field}_domain_service.go` | ドメインサービス |
| `*_test.go` | テスト |

### トランザクション管理

HTTPメソッドに応じて自動的にトランザクション制御を行う：

```go
// db_middleware.go
if method == POST || PUT || DELETE || PATCH {
    tx, _ := conn.BeginTxx(ctx, nil)  // トランザクション開始
    defer func() {
        if len(c.Errors) > 0 {
            tx.Rollback()  // エラー時はロールバック
        } else {
            tx.Commit()    // 成功時はコミット
        }
    }()
    ctx = context.WithValue(ctx, config.DBKey, tx)
} else {
    ctx = context.WithValue(ctx, config.DBKey, conn)  // 読み取り専用
}
```

### エラーハンドリング

`customerr`パッケージで型付きエラーを定義：

| エラー型 | HTTPステータス |
|---------|---------------|
| `BadRequestErr` | 400 |
| `UnauthorizedErr` | 401 |
| `ForbiddenErr` | 403 |
| `NotFoundErr` | 404 |
| `ConflictErr` | 409 |
| `InternalErr` | 500 |

`ErrorHandlingMiddleware`でエラー型に応じたレスポンスを返却。

### データベーススキーマ

```
users ─────┬──< skills >──── tags
           │
           ├──< careers
           │
           ├──< plans ────────┬──< plan_tags >──── tags
           │                  │
           │                  └──< contracts
           │                  │
           │                  └──< contract_requests
           │
           └──< mentor_recruitments ──< mentor_recruitment_tags >──── tags
                              │
                              └──< mentor_recruitment_proposals

categories ──< plans
           └──< mentor_recruitments
```

### 環境構築

1. `.env`ファイルを作成（`.env.example`を参照）
2. `docker-compose -f .docker/compose.yml up -d`
3. `make test-db-setup`
4. `make migrate-up`

開発用コマンドは`Makefile`を参照。
