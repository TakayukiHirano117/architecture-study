# プロジェクト概要

メンターと相談したいユーザーをつなぐプラットフォームです。ユーザー登録・ログイン、メンター募集、提案、プラン作成、契約リクエストなどを扱います。ユースケースやドメイン境界の詳細はルートの [README.md](README.md) を参照してください。

# 技術スタック

| 領域 | 主な技術 |
|------|----------|
| バックエンド | Go 1.24+（[go.mod](backend/go.mod) 準拠）、Gin、PostgreSQL、sqlx、golang-migrate |
| フロントエンド | Next.js 16、React 19、TypeScript、Tailwind CSS 4（[package.json](frontend/package.json)） |
| インフラ・開発 | Docker Compose（`backend/.docker/`）、GitHub Actions（`.github/workflows/`） |

アーキテクチャは **DDD + オニオンアーキテクチャ**。レイヤー配置・依存の向き・トランザクション方針は README の「このアプリの設計・アーキテクチャについて」を一次情報として参照してください。

# コーディング規約

- **条件分岐**: `if` のネストは原則禁止。早期 return・ガード節・小さな関数への切り出しなどで平坦に保つ。
- **バックエンド**: DDD の層（domain / app / infra）に沿って実装する。新規機能の手順は `.cursor/skills/ddd-feature-implementation/SKILL.md` があればそれに従う。
- **コメント**: コードを読めば意図が伝わる箇所にはコメントを書かない。仕様上の制約・非自明な理由・外部契約など、読み手がコードだけでは拾えない情報に限定する。
- **整形**: Go は `go fmt` / `goimports`、フロントは ESLint・プロジェクト既存のスタイルに合わせる。

# よく使うコマンド

- Docker 起動・停止: ルートで `make up` / `make down`（[Makefile](Makefile)）
- DB マイグレーション: `make migrate-up` など（同上）
- バックエンドテスト: `cd backend && go test ./...`（CI と同じ除外が必要なら README / `.github/workflows/ci.yml` の `go test` 行を参照）
- フロント開発: `cd frontend && npm run dev`

詳細な環境構築手順は README の「環境構築」を参照してください。
