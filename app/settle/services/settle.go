package services

import (
	"bet/core/balance"
	"bet/core/const/balance_code"
	"bet/core/const/status/order_bet_status"
	"bet/core/play"
	"bet/core/play/init_play"
	"bet/model"
	"bet/utils"
	"errors"
	"gorm.io/gorm"
)

type Settle struct {
	GameCode   int32
	GameName   string
	Issue      int64
	OpenResult interface{}
}

func (this *Settle) Calc() error {
	playPlayList := []string{}
	utils.DB.Table("sys_game_play").Where("game_code=? and status=1", this.GameCode).Group("play_type").Select("play_type").Find(&playPlayList)
	wons := play.Won{
		GameCode: this.GameCode,
		GameName: this.GameName,
	}
	wons.List = make([]play.WonItem, 0)
	for _, v := range playPlayList {
		play := init_play.PlayList[v]
		if play != nil {
			play.Won(this.OpenResult, "", &wons)
		}
	}
	if len(wons.List) == 0 {

		return errors.New("No data")
	}

	playS := []string{}
	for _, i := range wons.List {
		playS = append(playS, i.PlayCode)
	}
	openItems := []model.GameOpenItem{}
	utils.DB.Table("game_open_item").Find(&openItems, "game_code=? and play_code in ?", this.GameCode, playS)
	if len(openItems) == 0 {
		return errors.New("No data open item")
	}
	orderBets := []model.OrderBetSettle{}
	utils.DB.Table("order_bet_settle").Find(&orderBets, "game_code=? and issue_detail=?", this.GameCode, this.Issue)
	if len(orderBets) == 0 {
		return errors.New("No order")
	}
	wonOrder := []model.OrderBet{}
	noWonOrder := []model.OrderBet{}
	for i, v := range orderBets {
		for _, o := range openItems {
			if o.PlayCode == v.PlayCode {
				orderBets[i].Status = order_bet_status.Won
				orderBets[i].Rate = o.Rate
				orderBets[i].Won = o.Rate / 100 * v.Amount

			}
		}
		if v.Status == order_bet_status.Won {
			wonOrder = append(wonOrder, model.OrderBet(v))
		} else {
			v.Status = order_bet_status.Lost
			noWonOrder = append(noWonOrder, model.OrderBet(v))
		}

	}
	utils.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.CreateInBatches(noWonOrder, 100).Error
		for _, no := range noWonOrder {
			settle := model.OrderBetSettle{
				OrderNo: no.OrderNo,
			}
			if tx.Delete(settle).Error != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
		return nil
	})

	for _, i := range wonOrder {
		balanceUpdate := balance.BalanceUpdate{
			Amount:          i.Won,
			Uid:             i.Uid,
			BalanceCode:     balance_code.Won,
			OrderNoRelation: i.OrderNo,
			Title:           "Won",
			Note:            "Won " + i.Title,
		}
		balanceUpdate.Update(func(tx *gorm.DB) error {
			settle := model.OrderBetSettle{
				OrderNo: i.OrderNo,
			}
			err := tx.Delete(&settle).Error
			if err != nil {
				return err
			}
			err = tx.Create(i).Error
			if err != nil {
				return err
			}
			return nil
		})
	}
	return nil

}
