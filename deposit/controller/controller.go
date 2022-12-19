package controller

import (
	"bet/core/auth"
	services2 "bet/deposit/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func Deposit(gx *gin.Context) {
	uid := gx.Keys["Uid"]
	info := uid.(*auth.UserInfo)
	deposit := &services2.Deposit{}
	gx.ShouldBindBodyWith(deposit, binding.JSON)
	gx.JSON(http.StatusOK, deposit.DepositFunc(info))
}
