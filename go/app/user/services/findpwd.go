package services

import (
	"bet/db"
	"bet/model"
	"bet/utils"
	"fmt"
	"math/rand"
	"time"
)

type Findpwd struct {
	NewPwd    string `json:"newPwd"`
	NewPwd2   string `json:"newPwd2"`
	VerifCode int    `json:"verifCode"`
	Email     string `json:"email"`
}

// 发送验证码
func (this *Findpwd) Next1() error {
	user := model.UserBase{}
	user.UserBaseDB().First(&user, "email=?", this.Email)
	if user.Uid == 0 {
		return fmt.Errorf("Do not have this account")
	}
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 生成5位数字随机数
	min := 10000
	max := 99999
	randomNumber := rand.Intn(max-min+1) + min
	err := utils.SendEmail(this.Email, "Retrieve password", fmt.Sprintf("Retrieve password code:%d", randomNumber))
	if err != nil {
		return err
	}
	err = db.RedisSet(fmt.Sprintf("findpwd_%s", this.Email), randomNumber, time.Minute*15)
	return err
}

func (this *Findpwd) Next2() error {
	key := fmt.Sprintf("findpwd_s", this.Email)
	var code int
	err2 := db.RedisGet(key, &code)
	if err2 != nil {
		return err2
	}
	if code != this.VerifCode {
		return fmt.Errorf("Verification code error")
	}
	user := model.UserBase{}
	user.UserBaseDB().First(&user, "email=?", this.Email)
	if user.Uid == 0 {
		return fmt.Errorf("Do not have this account")
	}
	if this.NewPwd2 != this.NewPwd {
		return fmt.Errorf("")
	}
	user.Password = utils.Md5Encryption(this.NewPwd)
	err2 = db.GormDB.Table("user_base").Select("password").Updates(&user).Error
	if err2 != nil {
		return err2
	}
	return nil
}
