---
name: ddd-controller
description: Implements Controller and routing in DDD + Gin. Use when implementing HTTP handlers, request binding, or registering routes.
---

# DDD Controller

Controller・ルーティングの実装パターン。詳細は `.cursor/rules/infra-and-controller-layer.md` を参照。

## ディレクトリ構成

```
src/core/infra/controllers/
└── <action>_<entity>_controller.go
```

## Controller実装

- `NewXxxController()` で依存を注入する（Repository実装、DomainService等）
- `ShouldBindJSON` でリクエストをバインドする
- エラーは `ctx.Error(err)` に渡す（Middlewareがハンドリング）
- AppServiceを Controller 内でインスタンス化して実行する

## ルーティング登録

`src/core/cmd/main.go` にルートを追加する。
