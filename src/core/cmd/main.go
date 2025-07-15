package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/src/core/infra/controllers"
)

func main() {
	// コントローラーを初期化
	controller := controllers.NewController()

	// ルーターを作成
	mux := http.NewServeMux()

	// ルーティングを設定
	controller.SetupRoutes(mux)

	// サーバーを起動
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Health check available at: http://localhost:8080/health")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
