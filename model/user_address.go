package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *UserAddress) UserAddressDB() *gorm.DB {
	return utils.DB.Table("user_address")
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
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.UserAddressDB().First(this, id)
		if this.Id > 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
