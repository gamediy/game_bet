package services

import (
	"bet/model"
	"bet/utils"
)

type AmountList struct {
	AmountType string ` validate:"required" json:"amount_type"`
}

func (this AmountList) GetListByType(amount_type string) *map[string][]model.SysAmountItem {
	amountList := []model.SysAmount{}
	utils.DB.Find(&amountList, "type=?", amount_type)
	list := make(map[string][]model.SysAmountItem, len(amountList))
	if len(amountList) > 0 {
		for _, v := range amountList {
			amountItemList := []model.SysAmountItem{}
			utils.DB.Find(&amountItemList, "amount_id=? and status=1", v.Id)
			list[v.Type] = amountItemList
		}

	}
	return &list
}

func (this AmountList) GetList() []model.SysAmountItem {
	items := make([]model.SysAmountItem, 0)
	utils.DB.Find(&items, "type=? and status=1", this.AmountType)
	return items

}
