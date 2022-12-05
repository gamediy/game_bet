package services

import (
	"bet/core/game"
	"bet/utils"
)

func GetIssue(gameCode int32) *utils.Result[game.GameIssueRespone] {
	return game.GetIssue(gameCode)
}
