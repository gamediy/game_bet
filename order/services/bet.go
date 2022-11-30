package services

import "bet/core"

type Bet struct {
	GameCode int32 `json:"game_code"`
	PlayCode int32 `json:"play_code"`
	Amount   int64 `json:"order"`
}

func (this *Bet) BetFunc(userInfo *core.UserInfo) {

}
