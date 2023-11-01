package api

import (
	_const "ManagerApi/const"
	"ManagerApi/model"
	"ManagerApi/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupCoinRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/coin")
	v1Router.Use(gin.Logger())
	v1Router.POST("create", createCoin)
	v1Router.GET("list", coinList)
	v1Router.GET("detail", coinDetail)
	v1Router.POST("delete", deleteCoin)
}

func createCoin(ctx *gin.Context) {
	var coin model.Coin
	if err := ctx.ShouldBind(&coin); err != nil {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	service.GetDB().Create(&coin)
	fmt.Println(coin)
	SuccessResponse(ctx, coin)
}

func coinList(ctx *gin.Context) {
	var coins []model.Coin
	service.GetDB().Find(&coins)
	SuccessResponse(ctx, coins)
}

func coinDetail(ctx *gin.Context) {
	coinId := ctx.Query("coin_id")
	var coin model.Chain
	service.GetDB().Where("id", coinId).First(&coin)
	SuccessResponse(ctx, coin)
}

func deleteCoin(ctx *gin.Context) {
	coinId := ctx.PostForm("coin_id")
	fmt.Println(coinId)
	if len(coinId) < 1 {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	var chain model.Chain
	service.GetDB().Where("id", coinId).Delete(&chain)
	SuccessResponse(ctx, nil)
}
