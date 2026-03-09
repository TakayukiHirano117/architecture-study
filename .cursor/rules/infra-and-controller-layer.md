---
alwaysApply: true
---

# Infra層・Controller 実装パターン

## Controller

### 構造

```go
type CreateUserController struct {
    userRepo          userdm.UserRepository
    isExistByUserName userdm.IsExistByUserNameDomainService
    buildTags         tagdm.BuildTagsDomainService
}

func NewCreateUserController() *CreateUserController {
    return &CreateUserController{
        userRepo:          rdbimpl.NewUserRepositoryImpl(),
        isExistByUserName: userdm.NewIsExistByUserNameDomainService(rdbimpl.NewUserRepositoryImpl()),
        buildTags:         tagdm.NewBuildTagsDomainService(rdbimpl.NewTagRepositoryImpl()),
    }
}
```

### リクエスト処理

```go
func (c *CreateUserController) Exec(ctx *gin.Context) {
    var req userapp.CreateUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.Error(customerr.BadRequest(err.Error()))
        return
    }

    service := userapp.NewCreateUserAppService(c.userRepo, c.isExistByUserName, c.buildTags)
    if err := service.Exec(ctx, &req); err != nil {
        ctx.Error(err)
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "created"})
}
```

- `ctx.ShouldBindJSON` でリクエストをバインドし、エラー時は `ctx.Error()` に渡して return
- AppServiceを Controller 内でインスタンス化して実行する
- エラーは `ctx.Error(err)` に渡す（Middlewareがハンドリング）
- 成功時は `ctx.JSON()` でレスポンスを返す

## Repository実装

### 構造

```go
type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
    return &UserRepositoryImpl{}
}
```

### DB接続の取得

```go
func (r *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
    db := rdb.ExecFromCtx(ctx)  // TxがあればTx、なければDB
    _, err := db.ExecContext(ctx, `INSERT INTO users ...`, user.ID().String(), ...)
    return err
}
```

### 取得クエリとモデルマッピング

複数テーブルのJOINはモデル（DTO）を経由して Entity に変換する。

```go
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id shared.UUID) (*userdm.User, error) {
    db := rdb.ExecFromCtx(ctx)

    var models []models.UserDetailModel
    err := db.SelectContext(ctx, &models, `
        SELECT u.id, u.name, u.email, s.id as skill_id, ...
        FROM users u
        LEFT JOIN skills s ON u.id = s.user_id
        WHERE u.id = $1
    `, id.String())
    if err != nil {
        return nil, err
    }
    if len(models) == 0 {
        return nil, nil
    }

    // モデル → Entity 変換
    return mapToUserEntity(models)
}
```

### モデル（DTO）

SQL結果のマッピング専用。Infra層内のみで使用する。

```go
// src/core/infra/models/user_detail_model.go
type UserDetailModel struct {
    ID    string         `db:"id"`
    Name  string         `db:"name"`
    Email string         `db:"email"`
    // LEFT JOIN の結果は null になりうるので sql.NullString を使う
    SkillID sql.NullString `db:"skill_id"`
}
```

## Middleware

### トランザクション管理

`src/core/infra/middlewares/` に配置。Controller には記述しない。

```go
func DBMiddleware(db *sqlx.DB) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        if needsTransaction(ctx.Request.Method) {
            tx, _ := db.BeginTxx(ctx, nil)
            ctx.Set(rdb.TxKey, tx)
            ctx.Next()
            if len(ctx.Errors) > 0 {
                tx.Rollback()
            } else {
                tx.Commit()
            }
        } else {
            ctx.Set(rdb.DBKey, db)
            ctx.Next()
        }
    }
}
```

## パッケージ構成

```
src/core/infra/
├── controllers/
│   ├── create_user_controller.go
│   └── update_user_controller.go
├── middlewares/
│   ├── db_middleware.go
│   └── error_middleware.go
├── models/
│   ├── user_detail_model.go
│   └── plan_model.go
└── rdbimpl/
    ├── user_repository_impl.go
    └── plan_repository_impl.go
```

## ルーティング

`src/core/cmd/main.go` または専用のルーターファイルにまとめる。

```go
v1 := r.Group("/api/v1")
{
    users := v1.Group("/users")
    users.POST("", controllers.NewCreateUserController().Exec)
    users.PUT("/:id", controllers.NewUpdateUserController().Exec)
}
```
