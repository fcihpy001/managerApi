package model

import (
	"gorm.io/gorm"
	"time"
)

type CodeRequest struct {
	GroupName string  `form:"group_name" binding:"required"`
	Days      uint    `form:"days" binding:"required"`
	Count     uint    `form:"count" binding:"required"`
	NFTType   NFTType `form:"type" binding:"required"`
}

type ActiveCode struct {
	gorm.Model
	Code       string  `gorm:"code; primaryKey; type:char(10)"`
	GroupName  string  `gorm:"type:varchar(20)"`
	Status     int     `gorm:"status; size:1"`
	Address    string  `gorm:"address; type: varchar(50)"`
	NFTType    NFTType `gorm:"nft_type; type: varchar(10)"`
	Expiration time.Time
}
type NFTType string

const (
	//会员，伙伴、创世、
	NFTypeMember    NFTType = "member"
	NFTypeFellow    NFTType = "fellow"
	NFTypeCommunity NFTType = "community"
	NFTypeMedal     NFTType = "medal"
)

type NFT struct {
	gorm.Model
	ContractAddress string
	Type            NFTType
}
