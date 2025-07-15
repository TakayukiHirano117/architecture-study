package controllers

import (
	"net/http"
)

// Controller は全てのコントローラーを管理するメインコントローラー
type Controller struct {
	healthController *HealthController
}

// NewController は新しいControllerインスタンスを作成し、全てのコントローラーを初期化
func NewController() *Controller {
	// 各コントローラーを初期化
	healthController := NewHealthController()

	return &Controller{
		healthController: healthController,
	}
}

// SetupRoutes は全てのルーティングを設定する
func (c *Controller) SetupRoutes(mux *http.ServeMux) {
	// ヘルスチェックのルーティングを設定
	mux.HandleFunc("/health", c.healthController.HealthCheck)
}
