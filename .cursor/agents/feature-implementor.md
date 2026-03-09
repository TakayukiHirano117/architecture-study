---
name: feature-implementor
description: 新機能をDDD + オニオンアーキテクチャに従って実装するサブエージェント。Domain層→App層→Infra層→Controllerの順に実装する。
model: inherit
---

# Feature Implementor Agent

新機能を DDD + オニオンアーキテクチャに従って実装するサブエージェント。

以下のルールを参照しながら実装を進める：
- `ddd-and-onion-architecture.md` — アーキテクチャ原則・依存関係の方向
- `domain-layer.md` — Entity / ValueObject / Repository / DomainService の実装パターン
- `application-layer.md` — AppService / QueryService の実装パターン
- `infra-and-controller-layer.md` — Controller / Repository実装 / Model の実装パターン

---

## 実装順序

**必ずこの順序で実装する。内側の層から外側に向かって進める。**

```
1. Domain層（Entity, ValueObject, Repository Interface, DomainService）
2. App層（AppService, QueryService）
3. Infra層（Repository実装, Model）
4. Controller
5. ルーティング登録
```

---

## 実装フロー（各ステップ）

### Step 1: 実装前の調査

実装を始める前に必ず以下を確認する：

1. 同じドメインの既存ファイルを読む（命名・構造の踏襲）
2. 既存の類似機能のAppServiceを読む（処理フローの参考）
3. 既存のControllerを1つ読む（依存の注入方法の確認）

### Step 2: Domain層の実装

```
src/core/domain/<xxxdm>/
├── <entity>.go                          # Entity（必要な場合）
├── <value_object>.go                    # ValueObject（必要な場合）
├── <entity>_repository.go              # Repository Interface（新規ドメインの場合）
└── <purpose>_domain_service.go         # DomainService（必要な場合）
```

- Entity / ValueObject は `NewXxx()` と `NewXxxByVal()` の2コンストラクタを実装する
- Repository Interface には `//go:generate mockgen` コメントを付ける
- DomainService はインターフェースと実装を同一ファイルに定義する
- **実装後にモックを生成する**: `make gomock-generate-all`

### Step 3: App層の実装

```
src/core/app/<xxxapp>/
└── <action>_<entity>_app_service.go    # AppService
```

- 処理フロー: VO生成 → DomainService → Entity生成 → Repository保存
- エラーは `customerr` パッケージのカスタムエラーを使う
- Infra層には一切依存しない（Repository Interfaceのみ使う）

### Step 4: Infra層の実装

```
src/core/infra/rdbimpl/
└── <entity>_repository_impl.go         # Repository実装

src/core/infra/models/
└── <entity>_model.go                   # Model（JOINが必要な場合）
```

- `rdb.ExecFromCtx(ctx)` でDB接続を取得する
- 複雑なJOINはモデル（DTO）経由でEntityに変換する
- マイグレーションが必要な場合は `src/db/migrations/` に追加する

### Step 5: Controllerの実装

```
src/core/infra/controllers/
└── <action>_<entity>_controller.go
```

- `NewXxxController()` で依存を注入する
- `ShouldBindJSON` でリクエストをバインドする
- エラーは `ctx.Error(err)` に渡す（Middlewareがハンドリング）

### Step 6: ルーティング登録

`src/core/cmd/main.go` にルートを追加する。

---

## TDDフロー

各層の実装はテストを先に書く。

1. テストを書く
2. テストが落ちることを確認する（`make test-docker`）
3. テストが通る最低限の実装をする
4. テストが通ることを確認する（`make test-docker`）
5. リファクタリングする

テストの書き方は `test-implementor` を参照。

---

## チェックリスト

実装完了前に以下を確認する：

- [ ] Domain層のRepository InterfaceにGo generateコメントがあるか
- [ ] `make gomock-generate-all` でモックを再生成したか
- [ ] 依存の方向が守られているか（App層がInfra層に依存していないか）
- [ ] カスタムエラー（`customerr`）を適切に使っているか
- [ ] `make test-docker` が通るか
- [ ] `make lint-docker` が通るか
