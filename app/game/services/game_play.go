package services

import (
	"bet/db"
	"bet/model"
	"log"
)

type GamePlay struct {
	GameCode int32                `json:"game_code"`
	PlayItem []model.ConfGamePlay `json:"play_item"`
}

func (this *GamePlay) GetGamePlay() *GamePlay {

	sysGamePlayList := []model.ConfGamePlay{}
	err := db.GormDB.Find(&sysGamePlayList, "game_code=? and status=1", this.GameCode).Error
	if err != nil {
		log.Println(err)
	}
	this.PlayItem = sysGamePlayList
	return this

}
