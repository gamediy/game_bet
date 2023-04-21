package services

import (
	"bet/model"
	"bet/utils"
)

func GetGameCategory() []model.SysGameCategory {
	category := []model.SysGameCategory{}
	utils.DB.Table("sys_game_category").Find(&category, "status=1").Order("sort desc ")
	return category
}
