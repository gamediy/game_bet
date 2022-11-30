package model

import (
	"bet/utils"
	"gorm.io/gorm"
	"time"
)

func (this *UserBase) UserBaseDB() *gorm.DB {
	return utils.DB.Table("user_base")
}

type UserBase struct {
	CreatedAt time.Time //wrap
	UpdatedAt time.Time

	Uid          int64 `gorm:"primary_key"`
	Account      string
	Email        string
	Password     string
	Xid          string
	Ip           string
	Client_agent string
	Mobile       string
	Status       int32
	Level_bet    int32
	Level_pay    int32
	Level_agent  int32
	Pid          int64
	ParentPath   string
	Country      string
}

func (UserBase) TableName() string {
	return "user_base"
}
