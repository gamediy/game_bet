package services

import (
	"bet/model"
	"bet/utils"
)

type GameOpen struct {
	GameCode int32            `json:"game_code" validate:"required"`
	Status   int              `json:"status"`
	Limit    int              `json:"limit" validate:"required"`
	Index    int              `json:"index" validate:"required"`
	OpenList []model.GameOpen `json:"open_list"`
}

func (this *GameOpen) GetList() utils.Result[*GameOpen] {
	res := utils.Result[*GameOpen]{
		Code:      500,
		IsSuccess: false,
	}
	opens := []model.GameOpen{}
	err := utils.InputValidate(this, &res)
	if err != nil {
		return res
	}
	if this.Status < 9 {
		model.GameOpenDB().Limit(this.Limit).Offset((this.Index-1)*this.Limit).Find(&opens, "game_code=? and status=?", this.GameCode, this.Status)

	} else {
		model.GameOpenDB().Limit(this.Limit).Offset((this.Index-1)*this.Limit).Find(&opens, "game_code=?", this.GameCode)

	}
	this.OpenList = opens
	res.Data = this
	return res
}
