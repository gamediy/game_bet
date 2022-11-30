package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (this *OrderBalance) OrderBalanceDB() *gorm.DB {
	return utils.DB.Table("order_amount")
}

type OrderBalance struct {
	OrderNo         int64 `gorm:"primary_key"`
	Uid             int64
	Account         string
	Pid             int64
	AmountCode      int32
	Title           string
	BalanceBefore   int64
	BalanceAfter    int64
	CreatedAt       time.Time
	Note            string
	OrderNoRelation int64
	TramperNo       string
	ParentPath      string
	Balance         int64
}

func (OrderBalance) TableName() string {
	return "order_amount"
}

func (this *OrderBalance) GetByOrderNoCache(order_no int32) *OrderBalance {
	redisKey := fmt.Sprintf("order_balance:order_no:%d", order_no)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.OrderBalanceDB().First(this, order_no)
		if this.OrderNo >= 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
