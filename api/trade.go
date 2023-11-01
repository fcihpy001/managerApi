package api

import (
	_const "ManagerApi/const"
	"ManagerApi/middleware"
	"ManagerApi/model"
	"ManagerApi/service"
	"ManagerApi/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/gin-gonic/gin"
	"log"
	"math/big"
)

var ethClient = utils.GetEthClient()

func SetupTradeRouter(router *gin.Engine) {
	v1Router := router.Group("/v1/trade")
	v1Router.Use(gin.Logger())
	v1Router.Use(middleware.RecoveryMiddleware(), middleware.VerifyHeader())
	v1Router.POST("add", addTrade)
	v1Router.GET("list", tradeList)
	v1Router.GET("query", query)
}

func addTrade(ctx *gin.Context) {
	var trade model.Trade
	err := ctx.ShouldBind(&trade)
	if err != nil {
		ErrorResp(ctx, 402, _const.ErrorBodyMsg, nil)
		return
	}

	err = service.GetDB().Create(&trade).Error
	if err != nil {
		log.Printf("入库失败 %v", err)
	}
	log.Println("数据入库成功。。。")
	SuccessResp(ctx, "success", nil)
}

func tradeList(ctx *gin.Context) {
	var txs []model.Trade
	service.GetDB().Find(&txs)
	SuccessResponse(ctx, txs)

}

func query(ctx *gin.Context) {

	//获取前端转来的参数txhash
	txhash := ctx.Query("txhash")
	if len(txhash) < 32 {
		ErrorResp(ctx, 403, _const.ErrorBodyMsg, nil)
		return
	}

	//先从库里查
	var transaction model.Transaction
	service.GetDB().Where("tx_hash = ? AND `success = 1`", txhash).First(&transaction)
	if transaction.ID == 0 {
		//据交易hash 去查询交易详情
		transaction = queryByTxHash(txhash, utils.GetABI("./ABI/usdt.json"))
		//数据入库
		err := service.GetDB().Create(&transaction).Error
		if err != nil {
			log.Printf("入库失败 %v", err)
		}
		log.Println("交易hash数据入库成功：", transaction.TxHash)
	}

	data := model.TradeResp{
		Hash:    txhash,
		From:    transaction.From,
		To:      transaction.To,
		Value:   transaction.Value,
		Nonce:   transaction.Nonce,
		Success: transaction.Success,
	}
	SuccessResp(ctx, "", data)
}

func queryByTxHash(txHash string, tokenABI abi.ABI) model.Transaction {
	var transfer model.Transaction
	tx, isPending, err := ethClient.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}
	transfer.TxHash = txHash
	transfer.Success = !isPending

	//解析发送方地址
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	transfer.ChainId = utils.BigIntToUint(chainID)
	if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
		transfer.From = from.Hex()
	}

	data := tx.Data()
	methodID := data[:4] // 前四个字节是方法 ID

	// 找到匹配的 ABI 方法
	method, err := tokenABI.MethodById(methodID)
	if err != nil {
		log.Fatalf("无法找到匹配的 ABI 方法: %v", err)
	}
	transfer.Function = method.String()

	// 解码输入数据
	inputs, err := method.Inputs.UnpackValues(data[4:]) // 去除前四个字节
	if err != nil {
		log.Fatalf("无法解码输入数据: %v", err)
	}
	intVal := new(big.Int)
	intVal, success := intVal.SetString(fmt.Sprintf("%s", inputs[1]), 10)
	if !success {
		log.Fatalf("无法将字符串转换为 big.Int: %v", success)
	}

	// 计算value值 除以 10^18 并转换为 big.Float
	fltVal := new(big.Float).SetInt(intVal)
	scale := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	fltVal.Quo(fltVal, scale)

	transfer.To = fmt.Sprintf("%s", inputs[0])
	transfer.Value = fmt.Sprintf("%.18f", fltVal)
	transfer.Nonce = tx.Nonce()

	return transfer
}
