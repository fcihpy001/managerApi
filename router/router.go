package router

import (
	"ManagerApi/api"
	"ManagerApi/middleware"
	"ManagerApi/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	router := gin.Default()
	// 配置CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                                       // 允许所有来源，您可以根据需求设置特定的来源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许的HTTP方法
	router.Use(cors.New(config))

	setupUserRouter(router)

	setupActiveCodeRouter(router)

	err := router.Run(fmt.Sprintf(":%s", utils.Config.Server.Port))
	if err != nil {
		log.Println("服务器端口绑定失败", err)
	}
}

func setupUserRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/user")
	v1Router.Use(gin.Logger())
	v1Router.POST("login", api.Login)
	v1Router.POST("register", api.Register)
	v1Router.GET("info", middleware.AuthMiddleWare(), api.Info)
}

func setupActiveCodeRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/code")
	v1Router.Use(gin.Logger())
	v1Router.Use(middleware.CROSMiddleWare(), middleware.RecoveryMiddleware())
	v1Router.POST("create", api.Create)
	v1Router.GET("list", api.List)
	v1Router.POST("update", api.Update)
	v1Router.POST("enable", api.Enable)
}
