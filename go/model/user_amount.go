package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
)

func (this *UserAmount) UserAmountDB() *gorm.DB {
	return db.GormDB.Table("user_amount")
}

type UserAmount struct {
	Uid           int64 `gorm:"primary_key"`
	Email         string
	Balance       int64
	TotalBet      int64
	TotalDeposit  int64
	TotalWithdraw int64
	TotalProfit   int64
	TotalGift     int64
	Freeze        int64
	Account       string
	ParentPath    int32
}

func (UserAmount) TableName() string {
	return "user_amount"
}

func (this *UserAmount) GetByUidCache(uid int32) *UserAmount {
	redisKey := fmt.Sprintf("user_amount:uid:%d", uid)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.UserAmountDB().First(this, uid)
		if this.Uid >= 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
