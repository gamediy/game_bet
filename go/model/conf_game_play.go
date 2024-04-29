package model

import (
	"bet/db"
	"fmt"

	"gorm.io/gorm"
)

func (this *ConfGamePlay) ConfGamePlayDB() *gorm.DB {
	return db.GormDB.Table("conf_game_play")
}

type ConfGamePlay struct {
	Id       int32  `gorm:"primary_key" json:"id"`
	GameCode int32  `json:"game_code"`
	GameName string `json:"game_name"`
	Status   int32  `json:"status"`
	PlayCode string `json:"play_code"`
	PlayType string `json:"play_type"`
	Sort     int32  `json:"sort"`
	BetMin   int64  `json:"bet_min"`
	BetMax   int64  `json:"bet_max"`
}

func (ConfGamePlay) TableName() string {
	return "conf_game_play"
}

func (this *ConfGamePlay) GetByIdCache(id int32) *ConfGamePlay {
	redisKey := fmt.Sprintf("conf_game_play:id:%d", id)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.ConfGamePlayDB().First(this, id)
		if this.Id > 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
