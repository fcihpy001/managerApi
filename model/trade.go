package model

import "gorm.io/gorm"

type TradeResp struct {
	Hash    string `json:"hash"`
	From    string `json:"from"`
	To      string `json:"to"`
	Value   string `json:"value"`
	Nonce   uint64 `json:"nonce"`
	Success bool   `json:"success"`
}

// 交易记录
type Trade struct {
	gorm.Model
	From     string `json:"from" form:"from" binding:"required"`
	To       string `json:"to" form:"to" binding:"required"`
	ChainId  uint   `json:"chain_id" form:"chain_id" binding:"required"`
	CoinType string `json:"coin_type" form:"coin_type" binding:"required"`
	TXHash   string `json:"tx_hash" form:"tx_hash" binding:"required"`
	Type     string `json:"type" form:"type" binding:"required"`
	Remark   string `json:"remark" form:"remark"`
}

type TradeType string

const (
	TradeTypeBBS    TradeType = "bbs"
	TradeTypeCoupon TradeType = "coupon"
)
