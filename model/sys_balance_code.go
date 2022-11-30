package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *SysAmountCode) SysAmountCodeDB() *gorm.DB {
	return utils.DB.Table("blance_code")
}

type SysAmountCode struct {
	Code   int32 `gorm:"primary_key"`
	Remark string
	Status int32
	Type   string
	Title  string
}

func (SysAmountCode) TableName() string {
	return "blance_code"
}

func (this *SysAmountCode) GetByCodeCache(code int32) *SysAmountCode {
	redisKey := fmt.Sprintf("blance_code:code:%d", code)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysAmountCodeDB().First(this, code)
		if this.Code >= 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
