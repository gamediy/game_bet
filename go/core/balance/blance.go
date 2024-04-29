package balance

import (
	"bet/core/const/balance_code"
	"bet/db"
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

func (this *BalanceUpdate) Update(fc func(tx *gorm.DB) error) error {
	timeout, cancelFunc := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancelFunc()
	var err error

	for {
		err = this.updateExec(fc)
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			return nil
		}
		select {
		case <-timeout.Done():
			return err
		case <-time.After(1 * time.Second):

		}
	}
	return err
}
func (this *BalanceUpdate) updateExec(fc func(tx *gorm.DB) error) error {

	if this.Amount < 0 {
		this.Amount = -this.Amount
	}

	userBase := &model.UserBase{}
	userBase.UserBaseDB().First(userBase, this.Uid)
	userMoney := &model.UserAmount{}
	err := userMoney.UserAmountDB().Raw("select *  from user_amount where uid=? for update", this.Uid).Scan(userMoney).Error
	if err != nil {
		return err
	}
	if userMoney.Uid <= 0 {
		return fmt.Errorf("No such user")
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
	orderBalance.Balance = this.Amount
	orderBalance.Title = this.Title
	orderBalance.Pid = userBase.Pid
	orderBalance.BalanceCode = this.BalanceCode
	orderBalance.Note = this.Note
	orderBalance.ParentPath = userBase.ParentPath
	if this.BalanceCode <= 0 {
		if userMoney.Balance < this.Amount {

			return fmt.Errorf("Insufficient account balance")
		}
		orderBalance.Balance = -this.Amount
		orderBalance.BalanceAfter = userMoney.Balance - this.Amount
	}
	if orderBalance.BalanceAfter < 0 {

		return fmt.Errorf("Balance error")
	}

	switch this.BalanceCode {
	case balance_code.Bet:
		userMoney.TotalBet += this.Amount
		userMoney.TotalProfit -= this.Amount
	case balance_code.Deposit:
		userMoney.TotalDeposit += this.Amount
	case balance_code.Withdraw:
		userMoney.TotalWithdraw += this.Amount
	case balance_code.Won:
		userMoney.TotalProfit += this.Won

	}

	err = db.GormDB.Transaction(func(tx *gorm.DB) error {
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
			map[string]interface{}{
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

		return err
	}
	return nil
}
