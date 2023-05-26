package controller

import (
	services2 "bet/app/deposit/services"
	"bet/app/game/services"
	"bet/core/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Withdraw(gx *gin.Context) {

	uid := gx.Keys["Uid"]
	info := uid.(*auth.UserInfo)
	withdraw := &services.Withdraw{}
	gx.BindJSON(withdraw)
	gx.JSON(http.StatusOK, withdraw.WithdrawFunc(info))
}

func Deposit(gx *gin.Context) {
	uid := gx.Keys["Uid"]
	info := uid.(*auth.UserInfo)
	deposit := &services2.Deposit{}
	gx.BindJSON(deposit)
	gx.JSON(http.StatusOK, deposit.DepositFunc(info))
}
func Bet(gx *gin.Context) {
	bet := &services.Bet{}
	gx.BindJSON(bet)
	gx.JSON(http.StatusOK, bet.BetFunc(GetUserInfo(gx)))
}
func GetUserInfo(gx *gin.Context) *auth.UserInfo {
	uid := gx.Keys["Uid"]
	info := uid.(*auth.UserInfo)
	return info
}
