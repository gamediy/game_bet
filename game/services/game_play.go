package services

import (
	"bet/model"
	"bet/utils"
)

type GamePlay struct {
	GameCode int32               `json:"game_code"`
	PlayItem []model.SysGamePlay `json:"play_item"`
}

func (this *GamePlay) GetGamePlay() utils.Result[*GamePlay] {
	res := utils.Result[*GamePlay]{
		Code:      200,
		IsSuccess: true,
	}
	sysGamePlayList := []model.SysGamePlay{}
	utils.DB.Find(&sysGamePlayList, "game_code=? and status=1", this.GameCode)
	this.PlayItem = sysGamePlayList
	res.Data = this
	return res

}
