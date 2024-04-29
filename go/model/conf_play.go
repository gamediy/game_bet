package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
)

func (this *ConfPlay) ConfPlayDB() *gorm.DB {
	return db.GormDB.Table("sys_play")
}

type ConfPlay struct {
	Code     int32  `gorm:"primary_key" json:"code"`
	Name     string `json:"name"`
	Status   int32  `json:"status"`
	TypeName string `json:"type_name"`
	TypeCode int32  `json:"type_code"`
}

func (ConfPlay) TableName() string {
	return "conf_play"
}

func (this *ConfPlay) GetByCodeCache(code int32) *ConfPlay {
	redisKey := fmt.Sprintf("conf_play:code:%d", code)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.ConfPlayDB().First(this, code)
		if this.Code >= 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
