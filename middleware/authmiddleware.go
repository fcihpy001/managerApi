package middleware

import (
	"ManagerApi/service"
	"ManagerApi/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("X-Token")
		fmt.Print("请求token", tokenStr)

		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token缺失",
			})
			ctx.Abort()
			return
		}
		//截取Bearer 之后的数据
		tokenStr = tokenStr[7:]
		token, claims, err := utils.ParseToken(tokenStr)

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 402,
				"msg":  "token无效",
			})
			ctx.Abort()
			return
		}

		//	验证用户是否存在
		user := service.GetUserByUid(claims.UserId)
		if user.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 403,
				"msg":  "用户不存在",
			})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
