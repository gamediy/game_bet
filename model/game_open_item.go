package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *GameOpenItem) GameOpenItemDB() *gorm.DB {
	return utils.DB.Table("game_open_item")
}

type GameOpenItem struct {
	Id         int64  `gorm:"primary_key" json:"id"`
	GameOpenId int64  `json:"game_open_id"`
	PlayCode   int32  `json:"play_code"`
	PlayId     int32  `json:"play_id"`
	PlayName   string `json:"play_name"`
	Rate       int64  `json:"rate"`
	Sort       int32  `json:"sort"`
	Status     int32  `json:"status"`
	GameName   string `json:"game_name"`
	GameCode   int32  `json:"game_code"`
	PlayType   string `json:"play_type"`
}

func (GameOpenItem) TableName() string {
	return "game_open_item"
}

func (this *GameOpenItem) GetByIdCache(id int32) *GameOpenItem {
	redisKey := fmt.Sprintf("game_open_item:id:%d", id)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.GameOpenItemDB().First(this, id)
		if this.Id > 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
