package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code": 405,
					"msg":  fmt.Sprint(err),
					"data": nil,
				})
				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}
