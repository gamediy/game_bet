package controller

import (
	"bet/core/auth"
	"bet/game/services"
	"bet/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AmountList(gx *gin.Context) {

	value := gx.Query("amount_type")
	gx.JSON(http.StatusOK, services.AmountList{
		AmountType: value,
	}.GetList())

}

func Issue(gx *gin.Context) {

	vale, err := strconv.ParseInt(gx.Query("game_code"), 10, 64)
	if err != nil {
		result := utils.Result[string]{}
		gx.JSON(http.StatusOK, result.Error("Wrong game code"))
	}
	gx.JSON(http.StatusOK, services.GetIssue(int32(vale)))

}
func GameOpen(gx *gin.Context) {
	open := services.GameOpen{}
	gx.BindJSON(&open)
	gx.JSON(http.StatusOK, open.GetList)

}
func DepositRecord(gx *gin.Context) {
	drp := services.DepositRecordReq{}
	gx.BindJSON(&drp)
	gx.JSON(http.StatusOK, drp.GetList(auth.GetUserInfo(gx)))

}
func GameList(gx *gin.Context) {
	gl := services.GameList{}
	gx.JSON(http.StatusOK, gl.GetGameList)

}
