package middlewares

import (
	"context"
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/config"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func DBMiddleware(conn *sqlx.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		method := c.Request.Method
		ctx := c.Request.Context()

		if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete || method == http.MethodPatch {
			tx, err := conn.BeginTxx(ctx, nil)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			defer func() {
				if len(c.Errors) > 0 {
					_ = tx.Rollback()
				} else {
					_ = tx.Commit()
				}
			}()

			ctx = context.WithValue(ctx, config.DBKey, tx)
		} else {
			ctx = context.WithValue(ctx, config.DBKey, conn)
		}

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
