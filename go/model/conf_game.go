package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (this *ConfGame) ConfGameDB() *gorm.DB {
	return db.GormDB.Table("conf_game")
}

type ConfGame struct {
	Code            int32     `gorm:"primary_key" json:"code"`
	Name            string    `json:"name"`
	Status          int32     `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	StartTime       string    `json:"start_time"`
	EndTime         string    `json:"end_time"`
	TotalIssue      int32     `json:"total_issue"`
	IntervalSeconds int64     `json:"interval_seconds"`
	Sort            int32     `json:"sort"`
	CloseSeconds    int32     `json:"close_seconds"`
}

func (ConfGame) TableName() string {
	return "conf_game"
}

func (this *ConfGame) GetByCodeCache(code int32) *ConfGame {
	redisKey := fmt.Sprintf("conf_game:code:%d", code)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.ConfGameDB().First(this, code)
		if this.Code >= 0 {
			db.RedisSet(redisKey, this, 3600)
		}
	}
	return this
}
