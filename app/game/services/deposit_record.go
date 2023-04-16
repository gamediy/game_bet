package services

import (
	"bet/core/auth"
	"bet/model"
	"fmt"
)

type DepositRecordReq struct {
	PageIndex int
	PageSize  int
	Status    int32
}

type DepositRecordRes struct {
	Status     int32  `json:"status"`
	StatusText string `json:"status_text"`
	Title      string `json:"title"`
	CreatedAt  string `json:"created_at"`
	Amount     int64  `json:"amount"`
}

func (this *DepositRecordReq) GetList(userInfo *auth.UserInfo) []*DepositRecordRes {
	deposit := model.OrderDeposit{}
	var list []*model.OrderDeposit
	db := deposit.OrderDepositDB()
	db.Where("uid=?", userInfo.Uid)
	if this.Status > 0 {
		db.Where("status=?", this.Status)
	}
	db.Limit(this.PageSize).Offset((this.PageIndex - 1) * this.PageSize).Find(&list)
	res := []*DepositRecordRes{}
	for _, orderDeposit := range list {
		var statusStr string = ""
		if orderDeposit.Status == 1 {
			statusStr = "Processing"
		} else if orderDeposit.Status == 2 {
			statusStr = "Successful"
		} else if orderDeposit.Status == 3 {
			statusStr = "Fail"
		}
		recordRes := DepositRecordRes{
			Status:     orderDeposit.Status,
			StatusText: statusStr,
			Title:      fmt.Sprintf("%v(%v)", orderDeposit.Currency, orderDeposit.Protocol),
			CreatedAt:  orderDeposit.CreatedAt.Format("01-02 15:04:05"),
			Amount:     orderDeposit.Amount / 100,
		}
		res = append(res, &recordRes)
	}
	return res
}
