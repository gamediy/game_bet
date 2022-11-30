package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *SysAmountItem) SysAmountItemDB() *gorm.DB {
	return utils.DB.Table("sys_amount_item")
}

type SysAmountItem struct {
	Code     int32  `gorm:"primary_key" json:"code"`
	Title    string `json:"title"`
	Status   int32  `json:"status"`
	Detail   string `json:"detail"`
	AmountId int32  `json:"amount_id"`
	Net      string `json:"net"`
	Min      int64  `json:"min"`
	Max      int64  `json:"max"`
	Fee      int64  `json:"fee"`
	Type     string `json:"type"`
	Logo     string `json:"logo"`
	Sort     int32  `json:"sort"`
	Category string `json:"category"`
	Country  string `json:"country"`
	Protocol string `json:"protocol"`
}

func (SysAmountItem) TableName() string {
	return "sys_amount_item"
}

func (this *SysAmountItem) GetByCodeCache(code int32) *SysAmountItem {
	redisKey := fmt.Sprintf("sys_amount_item:code:%d", code)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysAmountItemDB().First(this, code)
		if this.Code > 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
