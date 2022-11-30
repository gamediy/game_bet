package services

import (
	"bet/model"
	"bet/utils"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	err      error
)

type Register struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `validate:"required"`
	Password2   string
	Ip          string
	ClientAgent string `json:"client_agent"`
}

func (this Register) Logic() *utils.Result[string] {
	var result = &utils.Result[string]{
		Code:      500,
		IsSuccess: false,
	}
	err := utils.InputValidate(this, result)
	if err != nil {
		return result
	}

	if this.Password != this.Password2 {
		result.Message = "Password is different"
		return result
	}
	var count int64

	var userBaseModel = &model.UserBase{}
	err = utils.DB.Model(userBaseModel).Where("email=? or account=?", this.Email, this.Email).Count(&count).Error
	if err != nil {
		fmt.Println(err)
	}
	if count > 0 {
		result.Message = "Already registered"
		return result
	}
	userBaseModel.Status = 1
	userBaseModel.Email = this.Email
	userBaseModel.Password = utils.Md5Encryption(this.Password)
	userBaseModel.Account = this.Email
	userBaseModel.Ip = this.Ip
	userBaseModel.Client_agent = this.ClientAgent
	userMoneyModel := &model.UserAmount{}
	userMoneyModel.Account = userBaseModel.Account
	userMoneyModel.Email = userBaseModel.Email
	userMoneyModel.Balance = 0.0
	err = utils.DB.Transaction(func(tx *gorm.DB) error {
		create := tx.Create(userBaseModel)

		if create.Error != nil {
			return err
		}
		userMoneyModel.Uid = userBaseModel.Uid
		err = tx.Create(userMoneyModel).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		result.Message = err.Error()
		return result
	}
	result.Code = 200
	result.IsSuccess = true
	return result
}
