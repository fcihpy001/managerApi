package router

import (
	"ManagerApi/api"
	"ManagerApi/middleware"
	"ManagerApi/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	router := gin.Default()
	router.Use(middleware.CROSMiddleWare())
	// 配置CORS中间件
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"*"}                                       // 允许所有来源，您可以根据需求设置特定的来源
	//config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许的HTTP方法
	//router.Use(cors.New(config))
	//router.Use(cors.Default())

	api.SetupDefiRouter(router)

	api.SetupUserRouter(router)

	api.SetupActiveCodeRouter(router)

	api.SetupTradeRouter(router)

	api.SetupContractRouter(router)

	api.SetupInviteRouter(router)

	api.SetupChainRouter(router)

	api.SetupCoinRouter(router)

	api.SetupWalletRouter(router)

	err := router.Run(fmt.Sprintf(":%s", utils.Config.Server.Port))
	if err != nil {
		log.Println("服务器端口绑定失败", err)
	}
}
