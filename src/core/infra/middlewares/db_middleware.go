package middlewares

import (
	"context"
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/config"
	"github.com/gin-gonic/gin"
)

type ctxKey string

const DBKey ctxKey = "db"

func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbConfig := config.NewDBConfig()
		method := c.Request.Method

		db, err := dbConfig.Connect()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		defer db.Close()

		ctx := c.Request.Context()

		if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete || method == http.MethodPatch {
			tx, err := db.BeginTxx(ctx, nil)
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

			ctx = context.WithValue(ctx, DBKey, tx)
		} else {
			ctx = context.WithValue(ctx, DBKey, db)
		}

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
