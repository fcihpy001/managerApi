package main

import (
	"ManagerApi/service"
	"ManagerApi/utils"
	"context"
	"os"
	"os/signal"
)

func main() {
	//=======================================================
	// 1. 读取配置文件
	utils.InitConfig()
	//=======================================================
	// 2. 初始化数据库
	ctx, cancel := context.WithCancel(context.Background())
	service.Init(ctx)

	//=======================================================

	//=======================================================
	// 4. gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Block until a signal is received.
	<-c
	cancel()
	//bot.StartBot(context.Background())
}
