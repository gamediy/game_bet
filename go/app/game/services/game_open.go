package services

import (
	"bet/model"
	"bet/utils"
)

type GameOpen struct {
	GameCode int32 `json:"game_code" validate:"required"`
	Status   int   `json:"status"`
	Limit    int   `json:"limit" validate:"required"`
	Index    int   `json:"index" validate:"required"`
}
type GameOpenResponse struct {
	Issue            int64  `json:"issue"`
	OpenResult       string `json:"open_result"`
	OpenResultDetail string `json:"open_result_detail"`
	OpenAt           string `json:"open_at"`
}

func (this *GameOpen) GetList() ([]GameOpenResponse, error) {
	opens := []model.GameOpen{}
	err := utils.InputValidate(this)
	if err != nil {
		return nil, err
	}
	if this.Status < 9 {
		model.GameOpenDB().Limit(this.Limit).Offset((this.Index-1)*this.Limit).Find(&opens, "game_code=? and status=?", this.GameCode, this.Status)
	} else {
		model.GameOpenDB().Limit(this.Limit).Offset((this.Index-1)*this.Limit).Find(&opens, "game_code=?", this.GameCode)
	}
	re := []GameOpenResponse{}
	for _, open := range opens {
		re = append(re, GameOpenResponse{
			Issue:            open.Issue,
			OpenAt:           open.OpenAt.Format("2006"),
			OpenResult:       open.OpenResult,
			OpenResultDetail: open.OpenResultDetail,
		})
	}
	return re, nil
}
