package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
)

func (this *ConfAmountItem) ConfAmountItemDB() *gorm.DB {
	return db.GormDB.Table("conf_amount_item")
}

type ConfAmountItem struct {
	Code     int32  `gorm:"primary_key" json:"code"`
	Title    string `json:"title"`
	Status   int32  `json:"status"`
	Detail   string `json:"detail"`
	AmountId int32  `json:"amount_id"`
	Net      string `json:"net"`
	Min      int64  `json:"min"`
	Max      int64  `json:"max"`
	Fee      int64  `json:"fee"`
	Type     string `json:"type"`
	Logo     string `json:"logo"`
	Sort     int32  `json:"sort"`
	Category string `json:"category"`
	Country  string `json:"country"`
	Currency string `json:"currency"`
	Protocol string `json:"protocol"`
}

func (ConfAmountItem) TableName() string {
	return "conf_amount_item"
}

func (this *ConfAmountItem) GetByCodeCache(code int32) *ConfAmountItem {
	redisKey := fmt.Sprintf("conf_amount_item:code:%d", code)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.ConfAmountItemDB().First(this, code)
		if this.Code > 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
