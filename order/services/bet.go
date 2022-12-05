package services

import (
	"bet/core/auth"
	"bet/core/balance"
	"bet/core/const/balance_code"
	"bet/core/const/status/order_bet_status"
	"bet/model"
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Bet struct {
	GameCode int32   `json:"game_code"`
	PlayCode int32   `json:"play_code"`
	Amount   float32 `json:"amount"`
}

func (this *Bet) BetFunc(userInfo *auth.UserInfo) utils.Result[string] {
	betMoney := int64(this.Amount)
	game := &model.SysGame{}
	game.GetByCodeCache(this.GameCode)
	result := utils.Result[string]{
		Code:      500,
		IsSuccess: false,
	}
	if game.Status == 0 {
		result.Message = "Wrong game code"
		return result
	}
	play := &model.SysGamePlay{}
	play.SysGamePlayDB().First(play, "game_code=? and play_code=?", this.GameCode, this.PlayCode)
	if play.Status == 0 {
		result.Message = "Wrong play code"
		return result
	}
	if play.BetMin < betMoney || play.BetMax > betMoney {
		result.Message = fmt.Sprintf("Wrong amount minimum:%d maximum:%d", play.BetMin, play.BetMax)
		return result
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
		GameType:   game.Type,
		Won:        0,
		PlayCode:   play.PlayCode,
		PlayName:   play.PlayName,
		Title:      fmt.Sprintf("%v %v", game.Name, play.PlayName),
		ParentPath: userInfo.ParentPath,
		CreateAt:   time.Now(),
	}
	balance := balance.BalanceUpdate{
		BalanceCode:     balance_code.Bet,
		Amount:          betMoney,
		Title:           "Bet",
		OrderNoRelation: orderBet.OrderNo,
		Uid:             userInfo.Uid,
		Note:            fmt.Sprintf("Bet %v %v", game.Name, play.PlayName),
	}
	res := balance.Update(func(tx *gorm.DB) error {
		err := tx.Create(orderBet).Error
		if err != nil {
			return err
		}
		return nil

	})
	if !res.IsSuccess {
		result.Message = res.Message
		return result
	}
	result.Code = 200
	result.Message = "Bet success"
	return result

}
