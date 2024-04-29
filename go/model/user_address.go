package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
)

func (this *UserAddress) UserAddressDB() *gorm.DB {
	return db.GormDB.Table("user_address")
}

type UserAddress struct {
	Id      int32  `gorm:"primary_key" json:"id"`
	Uid     int64  `json:"uid"`
	Account string `json:"account"`
	Address string `json:"address"`
	Net     string `json:"net"`
	Status  int32  `json:"status"`
}

func (UserAddress) TableName() string {
	return "user_address"
}

func (this *UserAddress) GetByIdCache(id int32) *UserAddress {
	redisKey := fmt.Sprintf("user_address:id:%d", id)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.UserAddressDB().First(this, id)
		if this.Id > 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
