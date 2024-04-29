package services

import (
	"bet/db"
	"bet/model"
)

func GetGameCategory() []model.ConfGameCategory {
	category := []model.ConfGameCategory{}
	db.GormDB.Table("conf_game_category").Find(&category, "status=1").Order("sort desc ")
	return category
}
