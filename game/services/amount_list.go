package services

import (
	"bet/model"
	"bet/utils"
)

type AmountList struct {
	AmountType string ` validate:"required"`
}

func (this AmountList) GetListByType(amount_type string) utils.Result[*map[string][]model.SysAmountItem] {
	amountList := []model.SysAmount{}
	utils.DB.Find(&amountList, "type=?", amount_type)
	list := make(map[string][]model.SysAmountItem, len(amountList))
	ajaxList := utils.Result[*map[string][]model.SysAmountItem]{}
	if len(amountList) > 0 {
		for _, v := range amountList {
			amountItemList := []model.SysAmountItem{}
			utils.DB.Find(&amountItemList, "amount_id=? and status=1", v.Id)
			list[v.Type] = amountItemList
		}

	}
	ajaxList.Code = 200
	ajaxList.IsSuccess = true
	ajaxList.Data = &list
	return ajaxList

}

func (this AmountList) GetList() utils.Result[[]model.SysAmountItem] {
	items := []model.SysAmountItem{}
	utils.DB.Find(&items, "type=? and status=1", this.AmountType)
	ajaxResult := utils.Result[[]model.SysAmountItem]{}
	ajaxResult.Data = items
	return ajaxResult

}
