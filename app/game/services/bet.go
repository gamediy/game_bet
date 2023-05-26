package services

import (
	"bet/core/auth"
	"bet/core/balance"
	"bet/core/const/balance_code"
	"bet/core/const/status/order_bet_status"
	game2 "bet/core/game"
	"bet/model"
	"bet/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Bet struct {
	GameCode int32   `json:"game_code"`
	PlayCode string  `json:"play_code"`
	Amount   float32 `json:"amount"`
	Issue    int64   `json:"issue"`
}

func (this *Bet) BetFunc(userInfo *auth.UserInfo) error {

	betMoney := int64(this.Amount)
	game := &model.ConfGame{}
	game.GetByCodeCache(this.GameCode)
	if game == nil {
		return errors.New("Wrong game")
	}
	if game.Status == 0 {
		return errors.New("Wrong game code")
	}
	play := &model.ConfGamePlay{}
	play.ConfGamePlayDB().First(play, "game_code=? and play_code=?", this.GameCode, this.PlayCode)
	if play.Status == 0 {

		return errors.New("Wrong play code")
	}
	if play.BetMin < betMoney || play.BetMax > betMoney {

		return fmt.Errorf("Wrong amount minimum:%d maximum:%d", play.BetMin, play.BetMax)
	}
	issue, err := game2.GetIssue(this.GameCode)
	if err != nil {
		return err
	}
	if issue.Issue != this.Issue {
		return errors.New("Issue number error")
	}
	if issue.Status != 1 {
		return errors.New("No more bets")
	}
	orderBet := &model.OrderBet{
		OrderNo:    utils.SnowflakeId(),
		Status:     order_bet_status.Bet,
		Uid:        userInfo.Uid,
		Pid:        userInfo.Pid,
		Account:    userInfo.Account,
		GameCode:   game.Code,
		GameName:   game.Name,
		Amount:     betMoney,
		Issue:      this.Issue,
		Won:        0,
		PlayCode:   play.PlayCode,
		PlayType:   play.PlayType,
		Title:      fmt.Sprintf("%v %v", game.Name, play.PlayCode),
		ParentPath: userInfo.ParentPath,
		CreateAt:   time.Now(),
	}
	balance := balance.BalanceUpdate{
		BalanceCode:     balance_code.Bet,
		Amount:          betMoney,
		Title:           "Bet",
		OrderNoRelation: orderBet.OrderNo,
		Uid:             userInfo.Uid,
		Note:            fmt.Sprintf("Bet %v %v", game.Name, play.PlayCode),
	}

	err = balance.Update(func(tx *gorm.DB) error {
		err := tx.Create(orderBet).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err

}
