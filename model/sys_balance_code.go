package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *SysBalanceCode) SysBalanceCodeDB() *gorm.DB {
	return utils.DB.Table("balance_code")
}

type SysBalanceCode struct {
	Code   int32 `gorm:"primary_key"`
	Remark string
	Status int32
	Type   string
	Title  string
}

func (SysBalanceCode) TableName() string {
	return "balance_code"
}

func (this *SysBalanceCode) GetByCodeCache(code int32) *SysBalanceCode {
	redisKey := fmt.Sprintf("balance_code:code:%d", code)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysBalanceCodeDB().First(this, code)
		if this.Code >= 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
