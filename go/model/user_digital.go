package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
)

func (this *UserDigital) UserDigitalDB() *gorm.DB {
	return db.GormDB.Table("user_digital")
}

type UserDigital struct {
	Id           int32  `gorm:"primary_key" json:"id"`
	Address      string `json:"address"`
	Net          string `json:"net"`
	Status       int32  `json:"status"`
	Count        int32  `json:"count"`
	PrivateKey   string `json:"private_key"`
	TotalDeposit int64  `json:"total_deposit"`
	Uid          int64  `json:"uid"`
	Account      string `json`
}

func (UserDigital) TableName() string {
	return "user_digital"
}

func (this *UserDigital) GetByIdCache(id int32) *UserDigital {
	redisKey := fmt.Sprintf("user_digital:id:%d", id)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.UserDigitalDB().First(this, id)
		if this.Id > 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
