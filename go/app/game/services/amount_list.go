package services

import (
	"bet/db"
	"bet/model"
)

type AmountList struct {
	AmountType string ` validate:"required" json:"amount_type"`
}

func (this AmountList) GetListByType(amount_type string) *map[string][]model.ConfAmountItem {
	amountList := []model.ConfAmount{}
	db.GormDB.Find(&amountList, "type=?", amount_type)
	list := make(map[string][]model.ConfAmountItem, len(amountList))
	if len(amountList) > 0 {
		for _, v := range amountList {
			amountItemList := []model.ConfAmountItem{}
			db.GormDB.Find(&amountItemList, "amount_id=? and status=1", v.Id)
			list[v.Type] = amountItemList
		}

	}
	return &list
}

func (this AmountList) GetList() []model.ConfAmountItem {
	items := make([]model.ConfAmountItem, 0)
	db.GormDB.Find(&items, "type=? and status=1", this.AmountType)
	return items

}
