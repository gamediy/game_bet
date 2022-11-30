package services

import (
	"bet/core"
	"bet/core/const/blance_code"
	"bet/core/const/status/withdraw_status"
	"bet/model"
	"bet/net/tron"
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Withdraw struct {
	Amount         float32 `validate:"required,min=1" json:"order"`
	Address        string  `validate:"required" json:"address"`
	AmountItemCode int32   `validate:"required" json:"amount_item_code"`
}

func (this *Withdraw) WithdrawFunc(userInfo *core.UserInfo) utils.Result[string] {
	result := utils.Result[string]{
		Code:      500,
		IsSuccess: false,
	}
	err := utils.InputValidate(this, &result)
	if err != nil {
		return result
	}
	withdrawMoney := int64(this.Amount * 100)

	item := &model.SysAmountItem{}
	amountItem := item.GetByCodeCache(this.AmountItemCode)
	if amountItem == nil || amountItem.Status == 0 {
		result.Message = "Parameter error"
		return result
	}

	if withdrawMoney > amountItem.Max || withdrawMoney < amountItem.Min {
		result.Message = fmt.Sprintf("Wrong order minimum:%d maximum:%d", amountItem.Min, amountItem.Max)
		return result
	}
	if amountItem.Fee > 0 {

		withdrawMoney = withdrawMoney - amountItem.Fee
		if withdrawMoney <= 0 {
			result.Message = "Wrong order"
			return result
		}
	}

	userMoney := &model.UserAmount{}
	userMoney.UserAmountDB().First(userMoney, "uid=?", userInfo.Uid)
	if userMoney.Uid < 1 {
		result.Message = "Wrong user"
		return result
	}
	_, err = tron.TronGrpcCLient.GetAccount(this.Address)
	if err != nil {
		result.Message = "Wrong address"
		return result
	}

	if withdrawMoney > userMoney.Balance {
		result.Message = "Wrong order"
		return result
	}
	orderNo := utils.SnowflakeId()
	update := core.BalanceUpdate{
		Uid:             userMoney.Uid,
		OrderNoRelation: orderNo,
		Title:           "Withdraw",
		BalanceCode:     blance_code.Withdraw,
	}
	update.Update(func(tx *gorm.DB) error {
		withdraw := &model.OrderWithdraw{
			OrderNo:      orderNo,
			Uid:          userMoney.Uid,
			Account:      userMoney.Account,
			Pid:          userInfo.Pid,
			Status:       withdraw_status.Processing,
			FinishAt:     time.Time{},
			StatusRemark: "Processing",
			Address:      this.Address,
			Fee:          amountItem.Fee,
			Amount:       withdrawMoney,
			Net:          amountItem.Net,
		}
		err = withdraw.OrderWithdrawDB().Create(withdraw).Error
		if err != nil {
			return err
		}
		return nil
	})
	result.Message = "Withdraw processing"
	result.Code = 200
	result.IsSuccess = true
	return result
}
