package model

import (
	"gorm.io/gorm"
	"time"
)

type CodeRequest struct {
	GroupName string `form:"groupname" binding:"required"`
	Days      uint   `form:"days" binding:"required"`
	Count     uint   `form:"count" binding:"required"`
	NFT       string `form:"nft" binding:"required"`
}

// 激活码
type ActiveCode struct {
	gorm.Model
	Code          string `gorm:"code; primaryKey; type:char(10)"`
	GroupName     string `gorm:"group_name; type:varchar(20)"`
	Status        int    `gorm:"status; size:1"`
	WalletAddress string `gorm:"wallet_address; type: varchar(50)"`
	NFT           string `gorm:"nft; type: varchar(10)"`
	Expiration    time.Time
}

type CodeResp struct {
}
