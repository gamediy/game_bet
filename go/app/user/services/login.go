package services

import (
	"bet/utils"
)

type Login struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func (this Login) Login() *utils.Result[string] {
	var result = &utils.Result[string]{
		Code:      500,
		IsSuccess: false,
	}
	err := utils.InputValidate(this)
	if err != nil {
		return nil
	}

	return result
}
