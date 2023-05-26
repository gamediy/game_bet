package model

import (
	"bet/db"
	"fmt"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type WithdrawConf struct {
	Fee float32 `json:"fee"`
	Min float32
	Max float32
}

func (this *Conf) ConfDB() *gorm.DB {
	return db.GormDB.Table("conf")
}

type Conf struct {
	Code   string `gorm:"primary_key"`
	Detail string
	Status int32
	Type   string
	Remark string
}

func (Conf) TableName() string {
	return "conf"
}

func GetConf[T any](code string, data *T) *T {
	conf := Conf{}
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
func (this *Conf) GetByCodeCache(code string) *Conf {
	redisKey := fmt.Sprintf("sys_conf:code:%d", code)
	err := db.RedisGet(redisKey, this)
	if err != nil {
		this.ConfDB().First(this, code)
		if this.Code != "" {
			db.RedisSet(redisKey, this, -1)
		}
	}
	return this
}
