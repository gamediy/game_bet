package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func GameOpenDB() *gorm.DB {
	return db.GormDB.Table("game_open")
}

type GameOpen struct {
	Id               int64     `gorm:"primary_key" json:"id"`
	Status           int32     `json:"status"`
	OpenResult       string    `json:"open_result"`
	OpenSource       string    `json:"open_source"`
	OpenAt           time.Time `json:"open_at"`
	Issue            int64     `json:"issue"`
	GameCode         int32     `json:"game_code"`
	GameName         string    `json:"game_name"`
	CloseAt          time.Time `json:"close_at"`
	OpenResultDetail string    `json:"open_result_detail"`
	WonResult        string    `json:"won_result"`
	IssueDetail      int64     `json:"issue_detail"`
}

func (GameOpen) TableName() string {
	return "game_open"
}

func (this *GameOpen) GetByIdCache(id int32) *GameOpen {
	redisKey := fmt.Sprintf("game_open:id:%d", id)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		GameOpenDB().First(this, id)
		if this.Id > 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}