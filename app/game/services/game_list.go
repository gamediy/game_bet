package services

import (
	"bet/core/game"
	"bet/model"
	"bet/utils"
)

type GameList struct {
	Category int `json:"category" validate:"required"`
}

func (this *GameList) Func() ([]game.GameIssueRespone, error) {
	list := []game.GameIssueRespone{}
	gameList := []model.SysGame{}
	issueList := []game.GameIssueRespone{}
	err := utils.InputValidate(this)
	if err != nil {
		return list, err
	}
	utils.DB.Table("sys_game").Where("category_id").Find(&gameList)
	for _, sysGame := range gameList {
		issue, _ := game.GetIssue(sysGame.Code)
		issueList = append(issueList, issue)
	}
	return issueList, nil
}
