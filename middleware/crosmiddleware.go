package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CROSMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//ctx.Writer.Header().Set("Access-Control-Max-Age", "600")
		//ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		//fmt.Println("sss", ctx.Request.Method)
		//fmt.Println("met", http.MethodOptions)
		//if ctx.Request.Method == http.MethodOptions {
		//	//ctx.AbortWithStatus(200)
		//	//ctx.JSON(200, nil)
		//	ctx.JSON(http.StatusOK, "ok!")
		//
		//}
		//
		//ctx.Next()
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,sign,x-api-key,source,timestamp")

			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				//log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()

	}
}
