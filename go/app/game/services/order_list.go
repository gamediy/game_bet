package services

import (
	model2 "bet/app/model"
	"bet/core/auth"
	"bet/core/const/status/order_bet_status"
	db2 "bet/db"
	"bet/model"
)

type OrderList struct {
	model2.Page
	GameCode int32 `json:"game_code"`
}
type OrderListRes struct {
	Uid       int64   `json:"uid"`
	StatusStr string  `json:"status_str"`
	Status    int32   `json:"status"`
	Amount    float64 `json:"amount"`
}

func (this *OrderList) Func(info auth.UserInfo) []OrderListRes {
	list := []model.OrderBet{}
	listRes := []OrderListRes{}
	db := db2.GormDB.Table("order_bet")
	db.Where("uid=?", info.Uid)
	if this.GameCode > 0 {
		db.Where("game_code=?", this.GameCode)
	}
	db.Limit(this.PageSize).Offset((this.PageIndex - 1) * this.PageSize).Find(&list)
	for _, bet := range list {
		model := OrderListRes{
			Uid:       bet.Uid,
			Status:    bet.Status,
			StatusStr: order_bet_status.GetStatusStr(bet.Status),
		}
		listRes = append(listRes, model)
	}
	return listRes
}
