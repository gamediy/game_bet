package services

import (
	"bet/core/game"
)

type Issue struct {
	GameCode int32 `json:"game_code"`
}

func (this *Issue) Get() (game.GameIssueRespone, error) {
	return game.GetIssue(this.GameCode)
}
