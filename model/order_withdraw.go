package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (this *OrderWithdraw) OrderWithdrawDB() *gorm.DB {
	return utils.DB.Table("order_withdraw")
}

type OrderWithdraw struct {
	OrderNo        int64     `gorm:"primary_key" json:"order_no"`
	Account        string    `json:"account"`
	Uid            int64     `json:"uid"`
	Pid            int64     `json:"pid"`
	Status         int32     `json:"status"`
	FinishAt       time.Time `json:"finish_at"`
	Detail         string    `json:"detail"`
	StatusRemark   string    `json:"status_remark"`
	Amount         int64     `json:"order"`
	SysRemark      string    `json:"sys_remark"`
	Address        string    `json:"address"`
	AmountFinally  int64     `json:"amount_finally"`
	Fee            int64     `json:"fee"`
	CreatedAt      time.Time `json:"created_at"`
	Net            string    `json:"net"`
	AmountItemCode int32     `json:"amount_item_code"`
	Currency       string    `json:"currency"`
	Protocol       string    `json:"protocol"`
}

func (OrderWithdraw) TableName() string {
	return "order_withdraw"
}

func (this *OrderWithdraw) GetByOrderNoCache(order_no int32) *OrderWithdraw {
	redisKey := fmt.Sprintf("order_withdraw:order_no:%d", order_no)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.OrderWithdrawDB().First(this, order_no)
		if this.OrderNo > 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
