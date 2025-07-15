package controllers

import (
	"net/http"
)

// HealthController はヘルスチェック関連の処理を担当
type HealthController struct{}

// NewHealthController は新しいHealthControllerインスタンスを作成
func NewHealthController() *HealthController {
	return &HealthController{}
}

// HealthCheck はヘルスチェックエンドポイントのハンドラー
func (hc *HealthController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy"}`))
}
