package api

import (
	_const "ManagerApi/const"
	"ManagerApi/model"
	"ManagerApi/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupChainRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/chain")
	v1Router.Use(gin.Logger())
	v1Router.POST("create", createChain)
	v1Router.GET("list", chainList)
	v1Router.POST("delete", deleteChain)
}

type chainRequest struct {
	ChainId   string
	ChainName string
	IsTest    bool
	RPC       string
	Websocket string
}

func createChain(ctx *gin.Context) {
	var chain chainRequest
	err := ctx.ShouldBind(&chain)
	if err != nil {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	service.GetDB().Save(&chain)
	SuccessResponse(ctx, chain)
}

func chainList(ctx *gin.Context) {
	var chains []model.Chain
	service.GetDB().Find(&chains)
	SuccessResponse(ctx, chains)
}

func deleteChain(ctx *gin.Context) {
	chainId := ctx.PostForm("chain_id")
	fmt.Println(chainId)
	if len(chainId) < 1 {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	var chain model.Chain
	service.GetDB().Where("chain_id", chainId).Delete(&chain)
	SuccessResponse(ctx, nil)
}
