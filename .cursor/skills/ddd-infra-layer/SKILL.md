---
name: ddd-infra-layer
description: Implements Infra layer (Repository implementation, Model/DTO) in DDD. Use when implementing database access, repository implementations, or SQL mapping.
---

# DDD Infra Layer

Infra層の実装パターン。詳細は `.cursor/rules/infra-and-controller-layer.md` を参照。

## ディレクトリ構成

```
src/core/infra/rdbimpl/
└── <entity>_repository_impl.go         # Repository実装

src/core/infra/models/
└── <entity>_model.go                   # Model（JOINが必要な場合）
```

## Repository実装

- **Domain interface に準拠**: Domain層で定義された Repository Interface を実装する
- **DB接続**: `rdb.ExecFromCtx(ctx)` で取得（TxがあればTx、なければDB）
- 複雑なJOINはモデル（DTO）経由でEntityに変換する

## Model（DTO）

- SQL結果のマッピング専用。Infra層内のみで使用
- LEFT JOIN の結果は null になりうるので `sql.NullString` 等を使う

## マイグレーション

スキーマ変更が必要な場合は `src/db/migrations/` に追加する。
