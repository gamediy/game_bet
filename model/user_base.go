package model

import (
	"time"
)

type UserBase struct {
	CreatedAt time.Time //wrap
	UpdatedAt time.Time

	Uid          int32
	Account      string
	Email        string
	Password     string
	Xid          string
	Ip           string
	Client_agent string
	Mobile       string
	Status       int
	Level_bet    int
	Level_pay    int
	Level_agent  int
	Pid          int
}

func (UserBase) TableName() string {
	return "user_base"
}
