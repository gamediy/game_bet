package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (this *SysGame) SysGameDB() *gorm.DB {
	return utils.DB.Table("sys_game")
}

type SysGame struct {
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

func (SysGame) TableName() string {
	return "sys_game"
}

func (this *SysGame) GetByCodeCache(code int32) *SysGame {
	redisKey := fmt.Sprintf("sys_game:code:%d", code)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysGameDB().First(this, code)
		if this.Code >= 0 {
			utils.RedisSet(redisKey, this, 3600)
		}
	}
	return this
}
