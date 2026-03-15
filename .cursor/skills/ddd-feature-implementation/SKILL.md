---
name: ddd-feature-implementation
description: Orchestrates DDD + Onion Architecture feature implementation. Use when implementing new features, adding use cases, or when the user mentions DDD, Onion Architecture, or layer-by-layer implementation.
---

# DDD Feature Implementation

DDD + オニオンアーキテクチャに従った新機能の実装フローをオーケストレートする。

## 実装順序

**必ずこの順序で実装する。内側の層から外側に向かって進める。**

```
1. Domain層（Entity, ValueObject, Repository Interface, DomainService）
2. Infra層（Repository実装, Model）
3. App層（AppService, QueryService）※Infraに依存
4. Controller
5. ルーティング登録
```

**レイヤー構成（外側→内側）**: Controller → App → Infra → Domain  
App層はInfra層より外側にあり、Usecase（AppService）はInfra層のRepository実装に依存する。

各層の実装時は対応する skill を参照する：
- Domain層 → `ddd-domain-layer`
- Infra層 → `ddd-infra-layer`
- App層 → `ddd-app-layer`
- Controller → `ddd-controller`

## 実装前の調査

実装を始める前に必ず以下を確認する：

1. 同じドメインの既存ファイルを読む（命名・構造の踏襲）
2. 既存の類似機能のAppServiceを読む（処理フローの参考）
3. 既存のControllerを1つ読む（依存の注入方法の確認）

## TDDフロー

各層の実装はテストを先に書く。

1. テストを書く
2. テストが落ちることを確認する（`make test-docker`）
3. テストが通る最低限の実装をする
4. テストが通ることを確認する（`make test-docker`）
5. リファクタリングする

テストの書き方は `test-implementor` を参照。

## チェックリスト

実装完了前に以下を確認する：

- [ ] Domain層のRepository InterfaceにGo generateコメントがあるか
- [ ] `make gomock-generate-all` でモックを再生成したか
- [ ] 依存の方向が守られているか（Controller → App → Infra → Domain）
- [ ] カスタムエラー（`customerr`）を適切に使っているか
- [ ] `make test-docker` が通るか
- [ ] `make lint-docker` が通るか

## 実装完了後の必須実行

**実装完了後、以下のコマンドを順に実行し、全て通ることを確認する：**

1. `make gomock-generate-all`
2. `make format-docker`
3. `make lint-fix-docker`
4. `make test-docker`
