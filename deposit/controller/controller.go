package controller

import (
	"bet/core/auth"
	services2 "bet/deposit/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Deposit(gx *gin.Context) {
	uid := gx.Keys["Uid"]
	info := uid.(*auth.UserInfo)
	deposit := &services2.Deposit{}
	gx.BindJSON(deposit)
	gx.JSON(http.StatusOK, deposit.DepositFunc(info))
}
