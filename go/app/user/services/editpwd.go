package services

import (
	"bet/core/auth"
	"bet/model"
	"bet/utils"
	"fmt"
)

type editpwd struct {
	NewPwd     string `json:"newPwd"`
	NewPwd2    string `json:"newPwd2"`
	CurrentPwd string `json:"currentPwd"`
}

func (this *editpwd) Func(info auth.UserInfo) error {
	var err error
	user := model.UserBase{}
	user.UserBaseDB().First(&user, info.Uid)
	if user.Uid == 0 {
		return fmt.Errorf("Account does not exist")
	}
	if this.NewPwd != this.NewPwd2 {
		return fmt.Errorf("The passwords do not match")
	}
	pwd := utils.Md5Encryption(this.CurrentPwd)
	if user.Password != pwd {
		return fmt.Errorf("Incorrect password")
	}
	user.Password = utils.Md5Encryption(this.NewPwd2)
	err = user.UserBaseDB().Updates(&user).Error
	if err != nil {
		return err
	}
	return err
}
