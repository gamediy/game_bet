package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
)

func (this *ConfBalanceCode) ConfBalanceCodeDB() *gorm.DB {
	return db.GormDB.Table("conf_balance_code")
}

type ConfBalanceCode struct {
	Code   int32 `gorm:"primary_key"`
	Remark string
	Status int32
	Type   string
	Title  string
}

func (ConfBalanceCode) TableName() string {
	return "conf_balance_code"
}

func (this *ConfBalanceCode) GetByCodeCache(code int32) *ConfBalanceCode {
	redisKey := fmt.Sprintf("conf_balance_code:code:%d", code)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.ConfBalanceCodeDB().First(this, code)
		if this.Code >= 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
