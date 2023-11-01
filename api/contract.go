package api

import (
	"ManagerApi/model"
	"ManagerApi/service"
	"github.com/gin-gonic/gin"
)

func SetupContractRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/contract")
	v1Router.Use(gin.Logger())
	v1Router.GET("list", contractList)
}

func contractList(ctx *gin.Context) {
	var contracts []model.Contract
	service.GetDB().Find(&contracts)
	SuccessResp(ctx, "success", contracts)
}
