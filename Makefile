.PHONY: dev build clean

# 開発サーバーを起動（air使用）
dev:
	air

# アプリケーションをビルド
build:
	go build -o bin/main ./src/core/cmd/main.go

# ビルドファイルをクリーンアップ
clean:
	rm -rf tmp/
	rm -rf bin/
	rm -f build-errors.log

# 依存関係をダウンロード
deps:
	go mod download

# テストを実行
test:
	go test ./...

# アプリケーションを直接実行
run:
	go run ./src/core/cmd/main.go
