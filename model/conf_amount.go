package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
)

func (this *ConfAmount) ConfAmountDB() *gorm.DB {
	return db.GormDB.Table("conf_amount")
}

type ConfAmount struct {
	Id       int32  `gorm:"primary_key" json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Status   int32  `json:"status"`
	Type     string `json:"type"`
}

func (ConfAmount) TableName() string {
	return "sys_amount"
}

func (this *ConfAmount) GetByIdCache(id int32) *ConfAmount {
	redisKey := fmt.Sprintf("conf_amount:id:%d", id)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.ConfAmountDB().First(this, id)
		if this.Id > 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
