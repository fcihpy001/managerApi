package middleware

import (
	_const "ManagerApi/const"
	"ManagerApi/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func VerifyHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取消息头中内容
		apiKey := ctx.GetHeader("x-api-key")
		sign := ctx.GetHeader("sign")
		timestamp := ctx.GetHeader("timestamp")

		if len(apiKey) < 6 || len(sign) < 10 || len(timestamp) < 8 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 402,
				"msg":  _const.ErrorHeaderMsg,
				"data": nil,
			})
			ctx.Abort()
			return
		}
		currentTime := time.Now().Unix() * 1000
		timestampInt, _ := strconv.Atoi(timestamp)
		fmt.Println("interval:", currentTime-int64(timestampInt))
		fmt.Println("head——time", os.Getenv("HEADER_TIME"))
		interval, _ := strconv.Atoi(os.Getenv("HEADER_TIME"))
		if currentTime-int64(timestampInt) > 60000*int64(interval) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 406,
				"msg":  "时间戳有误",
				"data": nil,
			})
			ctx.Abort()
			return
		}
		calculateSign := utils.HashStr(apiKey, "toplink1688@#", timestamp)
		if sign != calculateSign {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 405,
				"msg":  _const.ErrorBadRequest,
				"desc": fmt.Sprintf("正确签名是:%v", calculateSign),
				"data": nil,
			})
			ctx.Abort()
			return
		}
	}
}
