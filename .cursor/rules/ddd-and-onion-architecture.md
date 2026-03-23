---
alwaysApply: true
---

# DDD + オニオンアーキテクチャ原則

## アーキテクチャ概要

DDD（ドメイン駆動設計）とオニオンアーキテクチャを組み合わせて採用している。

```
src/core/
├── domain/   # ドメイン層（最内層）
├── app/      # アプリケーション層
└── infra/    # インフラ層（最外層）
```

## 依存関係の方向

**依存は必ず内側に向かう。外側の層から内側の層のみ参照可能。**

```
Infra層（Controller, Repository実装）
  ↓ 依存
App層（AppService, QueryService）
  ↓ 依存
Domain層（Entity, ValueObject, Repository Interface, DomainService）
```

- Domain層はどの層にも依存しない
- App層はDomain層にのみ依存する（Infra層には依存しない）
- Infra層はApp層・Domain層の両方に依存する

## 各層の責務

### Domain層（`src/core/domain/`）

ビジネスロジックの中心。フレームワーク・DBに一切依存しない。

- **Entity**: ビジネス上の概念を表すオブジェクト。ライフサイクルと同一性を持つ
- **Value Object**: 値そのものを表す不変オブジェクト。バリデーションを内包する
- **Repository Interface**: データ永続化の抽象。実装はInfra層に置く
- **Domain Service**: 単一のEntityに収まらないビジネスロジック

### App層（`src/core/app/`）

ユースケースの調整役。ビジネスロジック自体は持たず、Domain層を組み合わせてユースケースを実現する。

- **AppService（UseCase）**: 1つのユースケースに対応する処理フローを記述
- **QueryService**: 参照系の軽量クエリ（Repository直接利用）

### Infra層（`src/core/infra/`）

外部との接続。DBアクセス、HTTPリクエスト/レスポンスを扱う。

- **Controller**: HTTPリクエストの受付・レスポンス返却
- **Repository実装**: Repository Interfaceの実装（DB操作）
- **Middleware**: 横断的関心事（トランザクション管理など）
- **Model（DTO）**: SQL結果のマッピング用中間オブジェクト

## このプロジェクト固有の設計方針

### Domain Serviceの配置

Domain Serviceはインターフェースと実装の**両方をDomain層に配置**する。
（Repository Interfaceのみ抽象化し、Domain ServiceはDomain層で完結させる）

```
src/core/domain/userdm/
├── is_exist_by_user_id_domain_service.go      # インターフェース + 実装
├── user_repository.go                          # インターフェースのみ
└── ...
```

### トランザクション管理

**Middlewareが自動制御**する。Repositoryやアプリケーションコードでトランザクションを明示的に開始・コミット・ロールバックしない。

- POST / PUT / DELETE / PATCH → Middlewareがトランザクション開始
- GET → 読み取り専用（トランザクションなし）
- エラー時 → Middlewareが自動ロールバック
- 成功時 → Middlewareが自動コミット

Repository実装はContextからDB接続（またはTx）を取得する：

```go
db := rdb.ExecFromCtx(ctx)  // Txがあればそれを、なければ通常のDB接続を返す
```

### ドメインパッケージの命名規則

各ドメインはサフィックス `dm` を付けたパッケージ名にする。

```
userdm/                # ユーザードメイン
plandm/                # プランドメイン
contractdm/            # 契約ドメイン
mentor_recruitmentdm/  # メンター募集ドメイン
tagdm/                 # タグドメイン
categorydm/            # カテゴリドメイン
shared/                # ドメイン共通（UUID等）
```
