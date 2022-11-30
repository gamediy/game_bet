package controller

import (
	"bet/core"
	"bet/order/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Withdraw(gx *gin.Context) {

	uid := gx.Keys["Uid"]
	info := uid.(*core.UserInfo)
	withdraw := &services.Withdraw{}
	gx.BindJSON(withdraw)
	gx.JSON(http.StatusOK, withdraw.WithdrawFunc(info))
}

func Deposit(gx *gin.Context) {
	uid := gx.Keys["Uid"]
	info := uid.(*core.UserInfo)
	deposit := &services.Deposit{}
	gx.BindJSON(deposit)
	gx.JSON(http.StatusOK, deposit.DepositFunc(info))
}
func AmountList(gx *gin.Context) {

	value := gx.Query("amount_net")
	gx.JSON(http.StatusOK, services.AmountList{
		AmountType: value,
	}.GetList())

}
