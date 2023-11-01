package api

import (
	"ManagerApi/model"
	"ManagerApi/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"time"
)

func updateWallet(w model.Wallet) {
	var wallet model.Wallet
	service.GetDB().Where("address = ?", w.Address).Find(&wallet)
	if wallet.ID > 0 {
		wallet.Referrer = w.Referrer
		wallet.BindTime = time.Now().Format("2006-01-02 15:32:22")
		wallet.BindSource = w.BindSource
		wallet.TPLAmount = w.TPLAmount
		wallet.USDTAmount = w.USDTAmount
		tx := service.GetDB().Clauses(clause.OnConflict{UpdateAll: true}).Create(&wallet)
		if tx.Error != nil {
			log.Println("wallet update fail", tx.Error)
		}
	} else {
		tx := service.GetDB().Create(&w)
		if tx.Error != nil {
			log.Println("wallet update fail", tx.Error)
		}
	}
}

func response(ctx *gin.Context, status int, code int, msg string, data interface{}) {
	ctx.AbortWithStatusJSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ErrorResp(ctx *gin.Context, code int, msg string, data interface{}) {
	response(ctx, http.StatusOK, code, msg, data)
}

func ErrorResponse(ctx *gin.Context, code int, msg string) {
	response(ctx, http.StatusOK, code, msg, nil)
}

func SuccessResp(ctx *gin.Context, msg string, data interface{}) {
	response(ctx, http.StatusOK, 200, msg, data)
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	response(ctx, http.StatusOK, 200, "success", data)
}
