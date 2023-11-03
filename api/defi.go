package api

import (
	_const "ManagerApi/const"
	"ManagerApi/middleware"
	"ManagerApi/model"
	"ManagerApi/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SetupDefiRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/defi")
	v1Router.Use(gin.Logger())
	v1Router.GET("rank", middleware.VerifyHeader(), rank)
	v1Router.GET("info", middleware.VerifyHeader(), defiInfo)

	v1Router.POST("/nft/add", addNFT)
	v1Router.GET("/nft/list", nftList)

	v1Router.POST("/usdt/add", addUSDT)
	v1Router.GET("/usdt/list", usdtList)

	v1Router.POST("/deposit/add", addDeposit)
	v1Router.GET("/deposit/list", depositList)

	v1Router.GET("record", loginRecord)

	v1Router.POST("reward/add", addReward)
	v1Router.GET("reward/list", middleware.VerifyHeader(), rewardList)
}

func rank(ctx *gin.Context) {
	var request model.RankRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	var ranks []model.Rank
	var data model.PageData
	var count int64
	service.GetDB().Model(&model.Wallet{}).Count(&count)
	offset := (request.PageNum - 1) * request.PageSize
	var sql string
	if request.RankType == model.RankTypeInvite {
		sql = fmt.Sprintf("SELECT address , invite_amount amount, "+
			"'inivte' as type FROM wallet ORDER BY invite_amount DESC "+
			"LIMIT %d OFFSET %d", request.PageSize, offset)

	} else if request.RankType == model.RankTypeUSDT {
		sql = fmt.Sprintf("SELECT address , usdt_amount amount, "+
			"'usdt' as type FROM wallet ORDER BY usdt_amount DESC "+
			"LIMIT %d OFFSET %d", request.PageSize, offset)

	} else if request.RankType == model.RankTypeCommunity {
		sql = fmt.Sprintf("SELECT address , community_nft_amount amount, "+
			"'communityNFT' as type FROM wallet ORDER BY community_nft_amount DESC "+
			"LIMIT %d OFFSET %d", request.PageSize, offset)

	} else if request.RankType == model.RankTypeMedal {
		sql = fmt.Sprintf("SELECT address, medal_nft_amount amount, "+
			"'medaNFT' as type FROM wallet ORDER BY medal_nft_amount DESC "+
			"LIMIT %d OFFSET %d", request.PageSize, offset)
	}

	service.GetDB().Raw(sql).Scan(&ranks)

	data.PageSize = uint8(request.PageSize)
	data.CurrentPage = uint8(request.PageNum)
	data.Total = int(count)
	data.List = ranks
	SuccessResponse(ctx, data)
}

func addUSDT(ctx *gin.Context) {
	var usdt model.USDT
	err := ctx.ShouldBind(&usdt)
	if err != nil {
		return
	}
	service.GetDB().Save(&usdt)
	SuccessResponse(ctx, usdt)
}

func usdtList(ctx *gin.Context) {
	var list []model.USDT
	wallet := ctx.Query("wallet")
	service.GetDB().Find(&list)
	if len(wallet) > 0 {
		service.GetDB().Where("")
	}
	SuccessResponse(ctx, list)
}

func nftList(ctx *gin.Context) {
	var list []model.NFT
	to := ctx.Query("to")
	status := ctx.Query("status")
	status_int, _ := strconv.Atoi(status)
	status_int = 1
	nftType := ctx.Query("type")

	if len(to) > 0 && len(nftType) > 0 {
		service.GetDB().
			Where("`wallet` = ? AND `type` = ? AND `status` = ?", to, nftType, status_int).
			Find(&list)
	} else {
		service.GetDB().Find(&list)
	}
	SuccessResponse(ctx, list)
}

func addNFT(ctx *gin.Context) {
	var nft model.NFT
	err := ctx.ShouldBind(&nft)
	if err != nil {
		return
	}
	service.GetDB().Save(&nft)
	SuccessResponse(ctx, nft)
}

func depositList(ctx *gin.Context) {
	var request model.DepositRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	var data model.PageData
	var list []model.Deposit
	offset := (request.PageRequest.PageNum - 1) * request.PageRequest.PageSize
	service.GetDB().
		//Select("'from','amount','type','profit'").
		//Where("'from' = ? and type = ?", request.Wallet, request.Type).
		Find(&list).
		Limit(request.PageRequest.PageSize).
		Offset(offset)
	var total int64
	service.GetDB().
		Model(&model.Deposit{}).Count(&total)
	//Where("'from' = ? and type = ?", request.Wallet, request.Type).

	data.List = list
	data.Total = int(total)
	//data.TotalPage = data.Total / request.PageRequest.PageSize
	data.CurrentPage = uint8(request.PageRequest.PageNum)
	SuccessResponse(ctx, data)
}

func addDeposit(ctx *gin.Context) {
	var deposit model.Deposit
	err := ctx.ShouldBind(&deposit)
	if err != nil {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	service.GetDB().Save(&deposit)
	SuccessResponse(ctx, deposit)
}

func addReward(ctx *gin.Context) {
	var reward model.Reward
	err := ctx.ShouldBind(&reward)
	if err != nil {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	tx := service.GetDB().Create(&reward)
	if tx.Error != nil {
		ErrorResponse(ctx, 406, "数据处理异常")
		return
	}
	SuccessResponse(ctx, reward)
}

// 奖励列表
func rewardList(ctx *gin.Context) {
	var request model.ListRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		ErrorResponse(ctx, 403, _const.ErrorBodyMsg)
		return
	}
	var rewards []model.Reward
	offset := (request.PageRequest.PageNum - 1) * request.PageRequest.PageSize

	fmt.Println("wallet", request.Wallet)
	sql := fmt.Sprintf("SELECT * FROM reward ")
	if len(request.Wallet) > 0 {
		sql = fmt.Sprintf("SELECT * FROM reward WHERE `wallet`='%s' ", request.Wallet)
	}

	if len(request.StartDate) > 0 {
		sql += fmt.Sprintf(" AND created_at BETWEEN '%s' AND '%s 23:59:59' ORDER BY created_at DESC", request.StartDate, request.EndDate)
	}

	if request.PageRequest.PageSize > 0 {
		sql += fmt.Sprintf(" LIMIT %d OFFSET %d", request.PageRequest.PageSize, offset)
	}

	service.GetDB().Raw(sql).
		Scan(&rewards)

	SuccessResponse(ctx, rewards)
}

func defiInfo(ctx *gin.Context) {
	data := make(map[string]uint)
	var wallet_amount int64
	var member_amount int64
	var usdt_amount int64
	var airdrop_amount int64

	service.GetDB().Raw("SELECT COUNT(DISTINCT wallet)  FROM login_record").Scan(&wallet_amount)
	data["wallet_amount"] = uint(wallet_amount)

	service.GetDB().Raw("SELECT COUNT(*)  FROM wallet").Scan(&member_amount)
	data["member_amount"] = uint(member_amount)

	service.GetDB().Raw("SELECT SUM(usdt_amount)  FROM reward WHERE reward_type = 'usdt'").Scan(&usdt_amount)
	data["usdt_amount"] = uint(usdt_amount)

	service.GetDB().Raw("SELECT COUNT(*)  FROM reward WHERE reward_type != 'usdt'").Scan(&airdrop_amount)
	data["airdrop_amount"] = uint(airdrop_amount)
	SuccessResponse(ctx, data)
}
