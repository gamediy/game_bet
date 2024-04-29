package services

import (
	"bet/core/auth"
	"bet/core/balance"
	"bet/core/const/amount_net"
	"bet/core/const/balance_code"
	"bet/core/const/status/withdraw_status"
	"bet/model"
	"bet/net/tron"
	"bet/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Withdraw struct {
	Amount         float32 `validate:"required,min=1" json:"amount"`
	Address        string  `validate:"required" json:"address"`
	AmountItemCode int32   `validate:"required" json:"amount_item_code"`
}

func (this *Withdraw) WithdrawFunc(userInfo *auth.UserInfo) error {
	err := utils.InputValidate(this)
	if err != nil {
		return err
	}
	withdrawMoney := int64(this.Amount * 100)

	item := &model.ConfAmountItem{}
	amountItem := item.GetByCodeCache(this.AmountItemCode)
	if amountItem == nil || amountItem.Status == 0 {

		return errors.New("Parameter error")
	}

	if withdrawMoney > amountItem.Max || withdrawMoney < amountItem.Min {

		return fmt.Errorf("Wrong order minimum:%d maximum:%d", amountItem.Min, amountItem.Max)
	}
	if amountItem.Fee > 0 {

		withdrawMoney = withdrawMoney + amountItem.Fee
		if withdrawMoney <= 0 {

			return fmt.Errorf("Wrong order")
		}
	}

	userMoney := &model.UserAmount{}
	userMoney.UserAmountDB().First(userMoney, "uid=?", userInfo.Uid)
	if userMoney.Uid < 1 {

		return fmt.Errorf("Wrong user")
	}
	if amountItem.Net == amount_net.TRON {
		_, err = tron.TronGrpcCLient.GetAccount(this.Address)
		if err != nil {

			return fmt.Errorf("Wrong address")
		}
	}
	if withdrawMoney > userMoney.Balance {

		return fmt.Errorf("Wrong order")
	}
	orderNo := utils.SnowflakeId()
	update := balance.BalanceUpdate{
		Uid:             userMoney.Uid,
		OrderNoRelation: orderNo,
		Title:           "Withdraw",
		BalanceCode:     balance_code.Withdraw,
		Amount:          withdrawMoney,
	}
	err = update.Update(func(tx *gorm.DB) error {
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
			Amount:       withdrawMoney - amountItem.Fee,
			Net:          amountItem.Net,
			Currency:     amountItem.Currency,
			Protocol:     amountItem.Protocol,
		}
		err = withdraw.OrderWithdrawDB().Create(withdraw).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
