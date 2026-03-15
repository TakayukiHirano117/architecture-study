---
name: ddd-app-layer
description: Implements Application layer (AppService, UseCase) in DDD. Use when implementing use cases, application services, or orchestrating domain logic.
---

# DDD App Layer

App層の実装パターン。詳細は `.cursor/rules/application-layer.md` を参照。

## ディレクトリ構成

```
src/core/app/<xxxapp>/
└── <action>_<entity>_app_service.go    # AppService
```

## 処理フロー

```
VO生成 → DomainService呼び出し → Entity生成 → Repository保存
```

## Request 型

AppService と同じファイルまたは同じパッケージに定義する。

## カスタムエラー

`src/support/customerr/` を使う：

- バリデーション・重複 → `customerr.BadRequest()`
- リソースが存在しない → `customerr.NotFound()`
- 認証エラー → `customerr.Unauthorized()`

## 依存関係

Infra層のRepository実装に依存する（Repository InterfaceはDomain層で定義、実装はInfra層）。
