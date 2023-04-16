package model

import (
	"bet/utils"
	"fmt"
	"gorm.io/gorm"
)

func (this *SysGameCategory) SysGameCategoryDB() *gorm.DB {
	return utils.DB.Table("sys_game_category")
}

type SysGameCategory struct {
	Id       int32  `gorm:"primary_key" json:"id"`
	Logo     string `json:"logo"`
	Name     string `json:"name"`
	Status   int32  `json:"status"`
	ParentId int32  `json:"parent_id"`
	Sort     int32  `json:"sort"`
}

func (SysGameCategory) TableName() string {
	return "sys_game_category"
}

func (this *SysGameCategory) GetByIdCache(id int32) *SysGameCategory {
	redisKey := fmt.Sprintf("sys_game_category:id:%d", id)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysGameCategoryDB().First(this, id)
		if this.Id > 0 {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
