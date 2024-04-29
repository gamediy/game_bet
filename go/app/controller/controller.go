package controller

import (
	"bet/core/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Context  *gin.Context
	Orm      *gorm.DB
	Errors   error
	UserInfo *auth.UserInfo
}

func (e *Controller) GetUserInfo() *Controller {
	uid := e.Context.Keys["Uid"]
	if uid == nil {
		e.UserInfo = nil
		return e
	}
	e.UserInfo = uid.(*auth.UserInfo)

	return e
}
func (e *Controller) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {

		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}
func (e *Controller) Bind(d interface{}) *Controller {
	err := e.Context.BindJSON(d)
	if err != nil {
		e.AddError(err)
		return e
	}
	return e
}

// MakeContext 设置http上下文
func (e *Controller) MakeContext(c *gin.Context) *Controller {
	e.Context = c
	return e
}

type Result struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	IsSuccess bool        `json:"is_success"`
	Message   interface{} `json:"message"`
}

func (e *Controller) Ok(data interface{}, msg string) *Controller {
	result := Result{}
	result.Code = 200
	result.Data = data
	result.IsSuccess = true
	result.Message = msg
	e.Context.JSON(200, result)
	return e
}
func (e *Controller) Error(data interface{}, msg string) *Controller {
	result := Result{}
	result.Code = 500
	result.Data = data
	result.IsSuccess = false
	result.Message = msg
	e.Context.JSON(500, result)
	return e
}
func (e *Controller) Response(values ...interface{}) *Controller {

	// 获取参数个数
	n := len(values)
	lastArg := values[n-1]
	if err, ok := lastArg.(error); ok {
		if err != nil {
			result := Result{}
			result.Code = 500
			result.IsSuccess = false
			result.Message = err.Error()

			e.Context.AbortWithStatusJSON(200, result)
			return e
		}

	}
	result := Result{}
	result.Code = 200
	result.IsSuccess = true
	result.Data = values[0]

	e.Context.AbortWithStatusJSON(200, result)
	return e

}
