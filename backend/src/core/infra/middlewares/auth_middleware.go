package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/src/core/infra/shared"
	"github.com/TakayukiHirano117/architecture-study/src/support/customerr"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			_ = ctx.Error(customerr.Unauthorized("authorization header is required"))
			ctx.Abort()
			return
		}

		// Bearer トークン形式の検証
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			_ = ctx.Error(customerr.Unauthorized("invalid authorization header format"))
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := shared.ParseJWT(tokenString)
		if err != nil {
			_ = ctx.Error(customerr.Unauthorized("invalid or expired token"))
			ctx.Abort()
			return
		}

		// UserIDをContextに格納（今後の利用のため）
		ctx.Set("userID", claims.UserID)

		ctx.Next()
	}
}
