package services

import (
	"bet/core/auth"
	"bet/core/const/status/deposit_status"
	"bet/model"
	"bet/net/tron"
	"bet/utils"
	"fmt"
)

type Deposit struct {
	Amount         float32 `json:"amount"`
	AmountItemCode int32   `json:"amount_item_code"`
}

func (this *Deposit) DepositFunc(info *auth.UserInfo) utils.Result[interface{}] {
	item := &model.SysAmountItem{}
	amountItem := item.GetByCodeCache(this.AmountItemCode)
	result := utils.Result[interface{}]{
		Code:      500,
		IsSuccess: false,
	}
	if amountItem.Code <= 0 || amountItem.Status == 0 {
		result.Message = "Parameter error"
		return result
	}
	money := int64(this.Amount * 100)
	if money > amountItem.Max || money < amountItem.Min {
		result.Message = fmt.Sprintf("Wrong order minimum:%d maximum:%d", amountItem.Min, amountItem.Max)
		return result
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
	}

	response := make(map[string]string)
	response["net"] = amountItem.Net

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
			deposit.Address = addr
			digital.UserDigitalDB().Create(&digital)
		}
		response["address"] = digital.Address

	} else if amountItem.Net == "ETH" {

	}

	orderDeposit := &model.OrderDeposit{}
	err := orderDeposit.OrderDepositDB().Create(deposit).Error
	if err != nil {
		result.Message = err.Error()
		return result
	}

	result.Code = 200
	result.IsSuccess = true
	return result

}
