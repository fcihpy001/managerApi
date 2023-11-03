package api

import (
	_const "ManagerApi/const"
	"ManagerApi/middleware"
	"ManagerApi/model"
	"ManagerApi/service"
	"ManagerApi/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupWalletRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/wallet")
	v1Router.Use(gin.Logger())
	v1Router.GET("list", walletList)
	v1Router.GET("record", loginRecord)
	v1Router.GET("info", walletInfo)
	v1Router.GET("/withdraw/list", withdrawList)
	v1Router.POST("login", middleware.VerifyHeader(), walletLogin)

}

func walletLogin(ctx *gin.Context) {

	var record model.LoginRecord
	err := ctx.ShouldBind(&record)
	if err != nil {
		ErrorResp(ctx, 403, _const.ErrorBodyMsg, nil)
		return
	}
	//数据入库
	service.GetDB().Create(&record)
	SuccessResp(ctx, "success", nil)
}

func loginRecord(ctx *gin.Context) {
	var record []model.LoginRecord
	service.GetDB().Find(&record)
	SuccessResponse(ctx, &record)
}

func walletList(ctx *gin.Context) {
	var wallets []model.Wallet
	service.GetDB().Find(&wallets)
	SuccessResponse(ctx, &wallets)
}

//
//func addReward(ctx *gin.Context) {
//	var reward model.Reward
//	err := ctx.ShouldBind(&reward)
//	if err != nil {
//		return
//	}
//
//	service.GetDB().Save(&reward)
//	SuccessResponse(ctx, reward)
//}

func walletInfo(ctx *gin.Context) {
	wallet := ctx.Query("wallet")
	if len(wallet) < 40 {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	var walletInfo model.Wallet
	err := service.GetDB().Where("address = ?", wallet).First(&walletInfo).Error
	if err != nil {
		ErrorResp(ctx, 406, "数据查询错误", nil)
		return
	}
	SuccessResponse(ctx, walletInfo)
}

func withdrawList(ctx *gin.Context) {
	var request model.ListRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	var rewards []model.WithdrawResp
	offset := (request.PageRequest.PageNum - 1) * request.PageRequest.PageSize

	fmt.Println("wallet", request.Wallet)
	sql := fmt.Sprintf("SELECT * FROM token WHERE `from`='%s' AND `to`='%s' "+
		"AND `type`='usdt' AND `source`='reward' ORDER BY created_at DESC", utils.CouponAddress(), request.Wallet)

	if request.PageRequest.PageSize > 0 {
		sql += fmt.Sprintf(" LIMIT %d OFFSET %d", request.PageRequest.PageSize, offset)
	}
	service.GetDB().Raw(sql).
		Scan(&rewards)

	SuccessResponse(ctx, rewards)
}
