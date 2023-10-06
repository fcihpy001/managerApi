package middleware

import (
	"github.com/gin-gonic/gin"
)

func CROSMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		//if ctx.Request.Method == http.MethodOptions {
		//	ctx.AbortWithStatus(200)
		//} else {
		//	ctx.Next()
		//}
	}
}
