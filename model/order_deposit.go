package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (this *OrderDeposit) OrderDepositDB() *gorm.DB {
	return utils.DB.Table("order_deposit")
}

type OrderDeposit struct {
	OrderNo        int64     `json:"order_no"`
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
	Net            string    `json:"net"`
	AmountItemCode int32     `json:"amount_item_code"`
	Protocol       string    `json:"protocol"`
	Currency       string    `json:"currency"`
}

func (OrderDeposit) TableName() string {
	return "order_deposit"
}

func (this *OrderDeposit) GetByOrderNoCache(order_no int32) *OrderDeposit {
	redisKey := fmt.Sprintf("order_deposit:order_no:%d", order_no)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.OrderDepositDB().First(this, order_no)
		if this.OrderNo > 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
