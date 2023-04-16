package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *SysPlay) SysPlayDB() *gorm.DB {
	return utils.DB.Table("sys_play")
}

type SysPlay struct {
	Code     int32  `gorm:"primary_key" json:"code"`
	Name     string `json:"name"`
	Status   int32  `json:"status"`
	TypeName string `json:"type_name"`
	TypeCode int32  `json:"type_code"`
}

func (SysPlay) TableName() string {
	return "sys_play"
}

func (this *SysPlay) GetByCodeCache(code int32) *SysPlay {
	redisKey := fmt.Sprintf("sys_play:code:%d", code)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysPlayDB().First(this, code)
		if this.Code >= 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
