package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *SysGameType) SysGameTypeDB() *gorm.DB {
	return utils.DB.Table("sys_game_type")
}

type SysGameType struct {
	Code   string `gorm:"primary_key" json:"code"`
	Status int32  `json:"status"`
	Logo   string `json:"logo"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Srot   int32  `json:"srot"`
}

func (SysGameType) TableName() string {
	return "sys_game_type"
}

func (this *SysGameType) GetByCodeCache(code string) *SysGameType {
	redisKey := fmt.Sprintf("sys_game_type:code:%d", code)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysGameTypeDB().First(this, code)
		if this.Code != "" {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
