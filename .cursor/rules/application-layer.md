---
alwaysApply: true
---

# Application層 実装パターン

## AppService（UseCase）

### 構造

```go
type CreateUserAppService struct {
    userRepo             userdm.UserRepository
    isExistByUserName    userdm.IsExistByUserNameDomainService
    buildTags            tagdm.BuildTagsDomainService
}

func NewCreateUserAppService(
    userRepo userdm.UserRepository,
    isExistByUserName userdm.IsExistByUserNameDomainService,
    buildTags tagdm.BuildTagsDomainService,
) *CreateUserAppService {
    return &CreateUserAppService{
        userRepo:          userRepo,
        isExistByUserName: isExistByUserName,
        buildTags:         buildTags,
    }
}
```

### 処理フロー

```go
func (s *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
    // 1. Value Object 生成（バリデーション）
    name, err := userdm.NewUserName(req.Name)
    if err != nil {
        return customerr.BadRequest(err.Error())
    }

    // 2. Domain Service 呼び出し（存在チェック等）
    isExist, err := s.isExistByUserName.Exec(ctx, *name)
    if err != nil {
        return err
    }
    if isExist {
        return customerr.BadRequest("user name already exists")
    }

    // 3. タグ等の関連Entityの構築
    tags, err := s.buildTags.Exec(ctx, req.Tags)
    if err != nil {
        return err
    }

    // 4. Entity 生成
    user, err := userdm.NewUser(userdm.NewUserID(), *name, ...)
    if err != nil {
        return err
    }

    // 5. Repository 保存
    return s.userRepo.Store(ctx, user)
}
```

### Request 型

AppService と同じファイルまたは同じパッケージに定義する。

```go
type CreateUserRequest struct {
    Name     string
    Email    string
    Password string
    Tags     []TagRequest
}

type TagRequest struct {
    ID   string
    Name string
}
```

## カスタムエラーの使い方

`src/support/customerr/` に定義されているカスタムエラーを使う。

```go
customerr.BadRequest("user name already exists")   // 400
customerr.NotFound("user not found")               // 404
customerr.Unauthorized("unauthorized")             // 401
customerr.InternalServerError("internal error")   // 500
```

- バリデーションエラー・重複エラー → `customerr.BadRequest()`
- リソースが存在しない → `customerr.NotFound()`
- 認証エラー → `customerr.Unauthorized()`
- DB エラー等予期しないエラー → そのまま return（Middleware がハンドリング）

## QueryService

参照系で複数テーブルをまたぐ軽量クエリを担う。App層に配置する。

```go
// インターフェース
type IsExistPlanUserCombinationFromContractsQueryService interface {
    Exec(ctx context.Context, planID shared.UUID, userID shared.UUID) (bool, error)
}

// 実装（App層内）
type isExistPlanUserCombinationFromContractsQueryService struct {
    contractRepo contractdm.ContractRepository
}

func NewIsExistPlanUserCombinationFromContractsQueryService(
    contractRepo contractdm.ContractRepository,
) IsExistPlanUserCombinationFromContractsQueryService {
    return &isExistPlanUserCombinationFromContractsQueryService{contractRepo: contractRepo}
}
```

- 更新系は AppService を使う。QueryService は参照専用
- インターフェースと実装を同じファイルに配置する

## パッケージ構成

```
src/core/app/
├── userapp/
│   ├── create_user_app_service.go
│   └── update_user_app_service.go
├── planapp/
│   └── create_plan_app_service.go
└── query_service/
    └── is_exist_plan_user_combination_from_contracts_query_service.go
```
