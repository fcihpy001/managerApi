package model

import (
	"gorm.io/gorm"
	"time"
)

type CodeRequest struct {
	GroupName string `form:"group_name" binding:"required"`
	Days      uint   `form:"days" binding:"required"`
	Count     uint   `form:"count" binding:"required"`
}

type ActiveCode struct {
	gorm.Model
	Code       string `gorm:"code; primaryKey; type:char(6)"`
	GroupName  string `gorm:"type:varchar(20)"`
	Status     int    `gorm:"status; size:1"`
	Address    string `gorm:"address; type: varchar(50)"`
	Expiration time.Time
}
