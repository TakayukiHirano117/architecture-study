package middlewares

import (
	"fmt"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/support/custom_error"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var err error
				switch x := r.(type) {
				case string:
					err = errors.New(x)
				case error:
					err = x
				default:
					err = fmt.Errorf("unknown panic: %v", r)
				}

				appErr := custom_error.InternalWrapf(err, "panic occurred")

				ctx.Error(appErr)

				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
