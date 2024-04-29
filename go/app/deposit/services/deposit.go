package services

import (
	"bet/core/auth"
	"bet/core/const/status/deposit_status"
	"bet/model"
	"bet/net/tron"
	"bet/utils"
	"errors"
	"fmt"
)

type Deposit struct {
	Amount         float32 `json:"amount"`
	AmountItemCode int32   `json:"amount_item_code"`
}

func (this *Deposit) DepositFunc(info *auth.UserInfo) error {
	item := &model.ConfAmountItem{}
	amountItem := item.GetByCodeCache(this.AmountItemCode)
	if amountItem == nil {
		return errors.New("Parameter error")
	}
	if amountItem.Code <= 0 || amountItem.Status == 0 {

		return errors.New("Parameter error")
	}
	money := int64(this.Amount * 100)
	if this.Amount > float32(amountItem.Max) || this.Amount < float32(amountItem.Min) {
		return fmt.Errorf("Wrong deposit minimum:%d maximum:%d", amountItem.Min, amountItem.Max)

	}
	deposit := &model.OrderDeposit{
		OrderNo:        utils.SnowflakeId(),
		Uid:            info.Uid,
		Account:        info.Account,
		Status:         deposit_status.Processing,
		Pid:            info.Pid,
		Amount:         money,
		StatusRemark:   "Processing",
		Net:            amountItem.Net,
		AmountItemCode: amountItem.Code,
		Protocol:       amountItem.Protocol,
		ParentPath:     info.ParentPath,
		Currency:       amountItem.Currency,
	}

	response := make(map[string]string)
	response["net"] = amountItem.Net
	response["protocol"] = amountItem.Protocol
	if amountItem.Net == "TRON" {
		digital := &model.UserDigital{}
		digital.UserDigitalDB().First(digital, "uid=? and status=1 and net=?", info.Uid, amountItem.Net)
		if digital.Id == 0 {
			privateKey, addr := tron.TronGrpcCLient.GenerateKey()
			digital.PrivateKey = privateKey
			digital.Address = addr
			digital.Status = 1
			digital.Uid = info.Uid
			digital.Account = info.Account
			digital.Net = "TRON"
			digital.UserDigitalDB().Create(&digital)
		}
		deposit.Address = digital.Address
		response["address"] = digital.Address
	} else if amountItem.Net == "ETH" {
	}
	err := deposit.OrderDepositDB().Create(deposit).Error
	return err

}
