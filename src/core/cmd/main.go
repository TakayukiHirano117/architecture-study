package main

import (
	"fmt"
	"log"

	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/db"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
)

func main() {
	// データベース接続の設定
	dbConfig := db.NewConfig()
	dbConnection, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConnection.Close()

	// リポジトリの初期化
	_ = rdbimpl.NewUserRepositoryImpl(dbConnection) // 将来的にDIコンテナで使用

	// Ginエンジンの作成
	router := gin.Default()

	// コントローラーの初期化（将来的にはDIコンテナで管理）
	controller := controllers.NewController()

	// ルートの設定
	controller.SetupRoutes(router)

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Health check available at: http://localhost:8080/health")
	fmt.Printf("Database connected to: %s\n", dbConfig.DSN())

	// サーバーの起動
	router.Run(port)
}
