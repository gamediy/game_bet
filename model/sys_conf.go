package model

import (
	"bet/utils"
	"fmt"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type WithdrawConf struct {
	Fee float32 `json:"fee"`
	Min float32
	Max float32
}

func (this *SysConf) SysConfDB() *gorm.DB {
	return utils.DB.Table("sys_conf")
}

type SysConf struct {
	Code   string `gorm:"primary_key"`
	Detail string
	Status int32
	Type   string
	Remark string
}

func (SysConf) TableName() string {
	return "sys_conf"
}

func GetSysConf[T any](code string, data *T) *T {
	conf := SysConf{}
	cache := conf.GetByCodeCache(code)
	if cache.Code == "" {
		return nil
	}
	err := json.Unmarshal([]byte(cache.Detail), data)
	if err != nil {
		return nil
	}
	return data

}
func (this *SysConf) GetByCodeCache(code string) *SysConf {
	redisKey := fmt.Sprintf("sys_conf:code:%d", code)
	err := utils.RedisGet(redisKey, this)
	if err != nil {
		this.SysConfDB().First(this, code)
		if this.Code != "" {
			utils.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
