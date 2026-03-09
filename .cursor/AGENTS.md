# Project Instructions

## Overview

メンターとメンティーをつなぐメンタリングプラットフォームのバックエンドAPI。
メンティーはメンターのプランを選択して契約でき、メンターはメンティーに対してプランを提案・募集できる。

**主なユースケース:**
- ユーザー新規作成・更新
- メンター募集の作成
- プランの作成
- 契約リクエストの送信
- 契約の承認

## Tech Stack

| カテゴリ | 技術 |
|--------|------|
| 言語 | Go 1.24 |
| Webフレームワーク | Gin |
| DB | PostgreSQL 17.1-alpine |
| DB接続 | sqlx |
| テスト | testify / gomock |
| コンテナ | Docker / Docker Compose |
| マイグレーション | golang-migrate |
| ホットリロード | Air |
| Lint | golangci-lint |

## Architecture

DDD（ドメイン駆動設計）+ オニオンアーキテクチャを採用。

```
src/core/
├── domain/     # ドメイン層: Entity, Value Object, Repository Interface, Domain Service
├── app/        # アプリケーション層: UseCase（AppService）
└── infra/      # インフラ層: Controller, Repository実装, Middleware, Model(DTO)
```

- **Domain Service** はインターフェースと実装の両方をDomain層に配置する。
- **トランザクション管理** はMiddlewareで自動制御。POST/PUT/DELETE/PATCH はトランザクション開始、GET は読み取り専用。

## Coding Conventions

- if文のネストは行わない
- `utils`, `services` など意図が不明瞭な名前は使用しない。目的を表す名前をつける
- コードを読めばわかることはコメントに書かない
- DRY・YAGNI・SOLIDの各原則に従う

## Development Flow

テスト駆動開発（TDD）に基づく実装順序を守る。

1. テストを書く
2. テストが落ちることを確認する
3. テストが通る最低限の実装を行う
4. テストが通ることを確認する
5. リファクタリングする

## Directory Structure

```
src/
├── core/
│   ├── cmd/main.go              # エントリーポイント
│   ├── domain/                  # ドメイン層
│   │   ├── userdm/
│   │   ├── plandm/
│   │   ├── mentor_recruitmentdm/
│   │   ├── contractdm/
│   │   ├── contract_requestdm/
│   │   ├── categorydm/
│   │   ├── tagdm/
│   │   └── shared/
│   ├── app/                     # アプリケーション層
│   │   ├── userapp/
│   │   ├── planapp/
│   │   ├── mentorrecruitmentapp/
│   │   ├── contractapp/
│   │   ├── contractrequestapp/
│   │   └── query_service/
│   └── infra/                   # インフラ層
│       ├── controllers/
│       ├── middlewares/
│       ├── models/
│       └── rdbimpl/
├── db/
│   ├── init.sql
│   ├── migrations/
│   └── rdb_handler.go
└── support/
    ├── customerr/               # カスタムエラー定義
    ├── mock/                    # gomock生成ファイル
    └── testhelper/              # テスト用ヘルパー

config/
└── db.go                        # DB接続設定（環境変数から読み込み）
```

## Testing Policy

- **Domain層**: Entity, Value Object のユニットテスト
- **Infra層**: QueryService, DomainService, Repository の統合テスト（実DBを使用）
- **App層**: UseCase（AppService）のユニットテスト（gomockでモック）
- テーブル駆動テスト方式で記述する

