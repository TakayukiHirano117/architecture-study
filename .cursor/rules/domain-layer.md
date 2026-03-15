---
alwaysApply: true
---

# Domain層 実装パターン

## Entity

### コンストラクタの2パターン

```go
// 新規作成時（createdAt / updatedAt を自動設定）
func NewUser(id shared.UUID, name UserName, email Email, ...) (*User, error) {
    now := time.Now()
    return &User{
        id:        id,
        name:      name,
        email:     email,
        createdAt: now,
        updatedAt: now,
    }, nil
}

// DB復元時（時刻はそのまま受け取る）
func NewUserByVal(id shared.UUID, name UserName, email Email, ..., createdAt time.Time, updatedAt time.Time) (*User, error) {
    return &User{
        id:        id,
        name:      name,
        email:     email,
        createdAt: createdAt,
        updatedAt: updatedAt,
    }, nil
}
```

### フィールドとGetter

全フィールドをprivateにして、Getterで公開する。

```go
type User struct {
    id    shared.UUID
    name  UserName
    email Email
}

func (u *User) ID() shared.UUID { return u.id }
func (u *User) Name() UserName  { return u.name }
func (u *User) Email() Email    { return u.email }
```

### ビジネスロジック

Entityに属するビジネスロジック（プロフィール更新など）はEntityのメソッドとして実装する。

```go
func (u *User) UpdateProfile(name UserName, selfIntroduction SelfIntroduction, ...) {
    u.name = name
    u.selfIntroduction = selfIntroduction
    u.updatedAt = time.Now()
}
```

### IDの定義

各ドメインのIDは `shared.UUID` のエイリアスとして定義する。

```go
type UserID = shared.UUID

func NewUserID() UserID {
    return shared.NewUUID()
}
```

## Value Object

### 基本パターン

```go
type Email string

// 新規作成・バリデーション付き
func NewEmail(email string) (*Email, error) {
    if email == "" {
        return nil, errors.New("email is required")
    }
    if len(email) > 255 {
        return nil, errors.New("email must be 255 characters or less")
    }
    // フォーマット検証
    e := Email(email)
    return &e, nil
}

// DB復元用（軽量バリデーション）
func NewEmailByVal(email string) (*Email, error) {
    e := Email(email)
    return &e, nil
}
```

### バリデーションエラーのメッセージ

英語の小文字スネークケースで記述する。例: `"email is required"`, `"title must be 50 characters or less"`

### 複合 Value Object

他のVOやドメインモデルを内包できる。

```go
type Skill struct {
    skillID         SkillID
    tag             tagdm.Tag
    evaluation      Evaluation
    yearsOfExperience YearsOfExperience
}
```

## Repository Interface

```go
//go:generate mockgen -source=user_repository.go -destination=../../../support/mock/domain/userdm/user_repository_mock.go

type UserRepository interface {
    Store(ctx context.Context, user *User) error
    Update(ctx context.Context, user *User) error
    FindByName(ctx context.Context, name UserName) (*User, error)
    FindByID(ctx context.Context, id shared.UUID) (*User, error)
}
```

- ファイル先頭に `//go:generate mockgen` コメントを必ず付ける
- mock の生成先は `src/support/mock/domain/<パッケージ名>/` に統一する

## Domain Service

インターフェースと実装を**同一ファイルに**定義する。

```go
//go:generate mockgen -source=is_exist_by_user_id_domain_service.go -destination=../../../support/mock/domain/userdm/is_exist_by_user_id_domain_service_mock.go

type IsExistByUserIDDomainService interface {
    Exec(ctx context.Context, id shared.UUID) (bool, error)
}

type isExistByUserIDDomainService struct {
    userRepo UserRepository
}

func NewIsExistByUserIDDomainService(userRepo UserRepository) IsExistByUserIDDomainService {
    return &isExistByUserIDDomainService{userRepo: userRepo}
}

func (s *isExistByUserIDDomainService) Exec(ctx context.Context, id shared.UUID) (bool, error) {
    user, err := s.userRepo.FindByID(ctx, id)
    if err != nil {
        return false, err
    }
    return user != nil, nil
}
```

- インターフェース名: `IsExistByXxxDomainService` / `BuildXxxDomainService` など目的を表す名前
- 実装の struct は unexported（小文字始まり）
- `Exec` メソッドで処理を実行する規約
