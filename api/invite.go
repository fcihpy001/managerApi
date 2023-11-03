package api

import (
	_const "ManagerApi/const"
	"ManagerApi/model"
	"ManagerApi/service"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupInviteRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/invite")
	v1Router.Use(gin.Logger())
	v1Router.POST("create", createInvite)
	v1Router.GET("list", inviteList)
}

func inviteList(ctx *gin.Context) {
	addr := ctx.Query("wallet")
	if len(addr) < 40 {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	var models []model.Wallet
	service.GetDB().Where("referrer = ?", addr).Find(&models)

	SuccessResp(ctx, "success", models)
}

func createInvite(ctx *gin.Context) {

	addr := ctx.PostForm("wallet")
	referrer := ctx.PostForm("referrer")
	source := ctx.PostForm("source")
	//count := ctx.PostForm("tpl_amount")
	if len(addr) < 40 || len(referrer) < 40 {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	//c, _ := strconv.Atoi(count)
	var wallet model.Wallet
	wallet.Address = addr
	wallet.Referrer = referrer
	wallet.BindTime = time.Now().Format("2006-03-21 11:25:33")
	//wallet.TPLAmount = uint(c)
	wallet.BindSource = model.BindSourceType(source)

	updateWallet(wallet)
	SuccessResponse(ctx, nil)
}
