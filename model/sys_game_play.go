package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *SysGamePlay) SysGamePlayDB() *gorm.DB {
	return utils.DB.Table("sys_game_play")
}

type SysGamePlay struct {
	Id       int32  `gorm:"primary_key" json:"id"`
	GameCode int32  `json:"game_code"`
	GameName string `json:"game_name"`
	PlayName string `json:"play_name"`
	Status   int32  `json:"status"`
	PlayCode int32  `json:"play_code"`
	PlayType string `json:"play_type"`
	Sort     int32  `json:"sort"`
	BetMin   int64  `json:"bet_min"`
	BetMax   int64  `json:"bet_max"`
}

func (SysGamePlay) TableName() string {
	return "sys_game_play"
}

func (this *SysGamePlay) GetByIdCache(id int32) *SysGamePlay {
	redisKey := fmt.Sprintf("sys_game_play:id:%d", id)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysGamePlayDB().First(this, id)
		if this.Id >= 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
