package main

import (
	"ManagerApi/router"
	"ManagerApi/service"
	"ManagerApi/utils"
)

func main() {
	//=======================================================
	// 1. 读取配置文件
	utils.InitConfig()

	//=======================================================
	// 2. 初始化数据库
	service.InitDB()

	router.InitRouter()
}
