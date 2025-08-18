package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/src/support/custom_error"
	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case custom_error.AppError:
				log.Printf("ERROR: %+v", e.Trace())
				ctx.AbortWithStatusJSON(e.Code(), gin.H{
					"message": fmt.Sprintf("%d: %s", e.Code(), e.Msg()),
				})
			default:
				log.Printf("FATAL: %+v", e)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Fatal",
				})
			}
		}

	}
}
