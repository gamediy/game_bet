package model

import (
	"bet/db"
	"fmt"
	"gorm.io/gorm"
)

func (this *ConfGameCategory) ConfGameCategoryDB() *gorm.DB {
	return db.GormDB.Table("conf_game_category")
}

type ConfGameCategory struct {
	Id       int32  `gorm:"primary_key" json:"id"`
	Logo     string `json:"logo"`
	Name     string `json:"name"`
	Status   int32  `json:"status"`
	ParentId int32  `json:"parent_id"`
	Sort     int32  `json:"sort"`
}

func (ConfGameCategory) TableName() string {
	return "conf_game_category"
}

func (this *ConfGameCategory) GetByIdCache(id int32) *ConfGameCategory {
	redisKey := fmt.Sprintf("conf_game_category:id:%d", id)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.ConfGameCategoryDB().First(this, id)
		if this.Id > 0 {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
