---
name: ddd-domain-layer
description: Implements Domain layer (Entity, ValueObject, Repository interface, DomainService) in DDD. Use when implementing domain logic, entities, value objects, or repository interfaces.
---

# DDD Domain Layer

Domain層の実装パターン。詳細は `.cursor/rules/domain-layer.md` を参照。

## ディレクトリ構成

```
src/core/domain/<xxxdm>/
├── <entity>.go                          # Entity（必要な場合）
├── <value_object>.go                   # ValueObject（必要な場合）
├── <entity>_repository.go               # Repository Interface（新規ドメインの場合）
└── <purpose>_domain_service.go          # DomainService（必要な場合）
```

## 実装ルール

### Entity / ValueObject

- `NewXxx()` と `NewXxxByVal()` の2コンストラクタを実装する
- 全フィールドをprivate、Getterで公開
- ビジネスロジックはEntityのメソッドとして実装

### Repository Interface

- ファイル先頭に `//go:generate mockgen` コメントを必ず付ける
- mock の生成先は `src/support/mock/domain/<パッケージ名>/` に統一

### DomainService

- インターフェースと実装を同一ファイルに定義する
- 実装の struct は unexported（小文字始まり）
- `Exec` メソッドで処理を実行する規約

## 実装後

**必ずモックを生成する**: `make gomock-generate-all`
