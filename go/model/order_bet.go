package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (this *OrderBet) OrderBetDB() *gorm.DB {
	return db.GormDB.Table("order_bet")
}

type OrderBet struct {
	OrderNo      int64     `gorm:"primary_key" json:"order_no"`
	Uid          int64     `json:"uid"`
	Pid          int64     `json:"pid"`
	Account      string    `json:"account"`
	GameCode     int32     `json:"game_code"`
	GameCategory string    `json:"game_category"`
	Amount       int64     `json:"amount"`
	Status       int32     `json:"status"`
	GameName     string    `json:"game_name"`
	Won          int64     `json:"won"`
	PlayCode     string    `json:"play_code"`
	PlayType     string    `json:"play_type"`
	Title        string    `json:"title"`
	ParentPath   string    `json:"parent_path"`
	OpenResult   string    `json:"open_result"`
	CreateAt     time.Time `json:"create_at"`
	SettleAt     time.Time `json:"settle_at"`
	Rate         int64     `json:"rate"`
	BetContent   string    `json:"bet_content"`
	Issue        int64     `json:"issue"`
	GameOpenId   int64     `json:"game_open_id"`
}

func (OrderBet) TableName() string {
	return "order_bet"
}

func (this *OrderBet) GetByOrderNoCache(order_no int64) *OrderBet {
	redisKey := fmt.Sprintf("order_bet:order_no:%d", order_no)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.OrderBetDB().First(this, order_no)
		if this.OrderNo > 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
