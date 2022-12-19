package services

import (
	"bet/model"
	"bet/utils"
)

type GameList struct {
	GameCode string `json:"game_code"`
	GameName string `json:"game_name"`

	GameList []*GameList `json:"game_list"`
}

func (this *GameList) GetGameList() utils.Result[[]GameList] {
	u := utils.Result[[]GameList]{
		Code:      200,
		IsSuccess: true,
	}
	gameType := []model.SysGameType{}
	game := []model.SysGame{}
	utils.DB.Table("sys_game_type").Find(&gameType, "status=1").Order("sort desc ")
	list := []GameList{}
	for _, sysGameType := range gameType {
		utils.DB.Table("sys_game").Find(&game, "status=1 and type=?", sysGameType.Code)
		gameList := GameList{
			GameCode: sysGameType.Code,
			GameName: sysGameType.Name,
		}
		for _, sysGame := range game {
			g := GameList{
				GameCode: string(sysGame.Code),
				GameName: sysGame.Name,
			}
			gameList.GameList = append(gameList.GameList, &g)
		}
		list = append(list, gameList)
	}
	u.Data = list
	return u
}
