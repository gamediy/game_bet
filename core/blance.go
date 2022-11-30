package core

import (
	"bet/core/const/blance_code"
	"bet/model"
	"bet/utils"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type BalanceUpdate struct {
	Uid             int64
	Amount          int64
	Won             int64
	OrderNoRelation int64
	BalanceCode     int32
	Title           string
	Note            string
}

func (this *BalanceUpdate) Update(fc func(tx *gorm.DB) error) *utils.Result[string] {
	timeout, cancelFunc := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancelFunc()
	var err error
	var res *utils.Result[string]
	for {
		res, err = this.updateExec(fc)
		fmt.Println(err.Error())
		if err == nil {
			return res
		}
		select {
		case <-timeout.Done():
			return res
		case <-time.After(1 * time.Second):

		}
	}
	return res
}
func (this *BalanceUpdate) updateExec(fc func(tx *gorm.DB) error) (*utils.Result[string], error) {

	if this.Amount < 0 {
		this.Amount = -this.Amount
	}
	result := &utils.Result[string]{
		Code:      500,
		IsSuccess: false,
	}

	userBase := &model.UserBase{}
	userBase.UserBaseDB().First(userBase, this.Uid)
	userMoney := &model.UserAmount{}
	err := userMoney.UserAmountDB().Raw("select *  from user_amount where uid=? for update", this.Uid).Scan(userMoney).Error
	if err != nil {
		result.Message = err.Error()
		return result, err
	}

	if userMoney.Uid <= 0 {
		result.Message = "No such user"
		return result, nil
	}
	orderBalance := &model.OrderBalance{}
	orderBalance.OrderNo = utils.SnowflakeId()
	orderBalance.OrderNoRelation = this.OrderNoRelation
	orderBalance.Uid = userMoney.Uid
	orderBalance.Pid = userBase.Pid
	orderBalance.Account = userMoney.Account
	orderBalance.ParentPath = userBase.ParentPath
	orderBalance.BalanceBefore = userMoney.Balance
	orderBalance.BalanceAfter = userMoney.Balance + this.Amount
	orderBalance.Title = this.Title
	orderBalance.Note = this.Note
	orderBalance.ParentPath = userBase.ParentPath
	if this.BalanceCode <= 0 {
		if userMoney.Balance < this.Amount {
			result.Message = "Insufficient account balance"
			return result, nil
		}
		orderBalance.Balance = -this.Amount
		orderBalance.BalanceAfter = userMoney.Balance - this.Amount
	}
	if orderBalance.BalanceAfter < 0 {
		result.Message = "Balance error"
		return result, nil
	}

	switch this.BalanceCode {
	case blance_code.Bet:
		userMoney.TotalBet += this.Amount
		userMoney.TotalProfit -= this.Amount
	case blance_code.Deposit:
		userMoney.TotalDeposit += this.Amount
	case blance_code.Withdraw:
		userMoney.TotalWithdraw += this.Amount
	case blance_code.Won:
		userMoney.TotalProfit += this.Won

	}

	err = utils.DB.Transaction(func(tx *gorm.DB) error {
		if fc != nil {
			err := fc(tx)
			if err != nil {
				return err
			}
		}
		err := tx.Create(orderBalance).Error
		if err != nil {
			return err
		}
		err = tx.Table("user_amount").Where("uid=? and balance=?", this.Uid, orderBalance.BalanceBefore).Updates(
			map[string]int64{
				"balance":        orderBalance.BalanceAfter,
				"total_bet":      userMoney.TotalBet,
				"total_deposit":  userMoney.TotalDeposit,
				"total_profit":   userMoney.TotalProfit,
				"total_withdraw": userMoney.TotalWithdraw,
			},
		).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		result.Message = err.Error()
		return result, err
	}
	result.Message = "Success"
	result.Code = 200
	return result, nil
}
