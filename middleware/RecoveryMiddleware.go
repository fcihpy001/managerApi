package middleware

import (
	"ManagerApi/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				api.ErrorResp(ctx, 405, fmt.Sprint(err), nil)
			}
		}()

		ctx.Next()
	}
}
