package services

import (
	"bet/model"
	"bet/utils"
	"log"
)

type GamePlay struct {
	GameCode int32               `json:"game_code"`
	PlayItem []model.SysGamePlay `json:"play_item"`
}

func (this *GamePlay) GetGamePlay() *GamePlay {

	sysGamePlayList := []model.SysGamePlay{}
	err := utils.DB.Find(&sysGamePlayList, "game_code=? and status=1", this.GameCode).Error
	if err != nil {
		log.Println(err)
	}
	this.PlayItem = sysGamePlayList
	return this

}
