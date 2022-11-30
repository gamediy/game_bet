package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *SysAmount) SysAmountDB() *gorm.DB {
	return utils.DB.Table("sys_amount")
}

type SysAmount struct {
	Id       int32  `gorm:"primary_key" json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Status   int32  `json:"status"`
	Type     string `json:"type"`
}

func (SysAmount) TableName() string {
	return "sys_amount"
}

func (this *SysAmount) GetByIdCache(id int32) *SysAmount {
	redisKey := fmt.Sprintf("sys_amount:id:%d", id)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysAmountDB().First(this, id)
		if this.Id > 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
