package services

import (
	"bet/model"
	"bet/utils"
)

func GetGameCategory() utils.Result[[]model.SysGameCategory] {
	u := utils.Result[[]model.SysGameCategory]{
		Code:      200,
		IsSuccess: true,
	}
	category := []model.SysGameCategory{}
	utils.DB.Table("sys_game_category").Find(&category, "status=1").Order("sort desc ")

	u.Data = category
	return u
}
