package controller

import (
	"bet/app/controller"
	services2 "bet/app/game/services"
	"bet/core/auth"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	controller.Controller
}

func (app Controller) AmountList(gx *gin.Context) {
	s := services2.AmountList{}
	app.MakeContext(gx).Bind(&s)
	app.Ok(s.GetList(), "")

}

func (app Controller) Issue(gx *gin.Context) {
	s := services2.Issue{}
	app.MakeContext(gx).Bind(&s)

	app.Response(s.Get())

}

func (app Controller) GameOpen(gx *gin.Context) {
	open := services2.GameOpen{}
	app.MakeContext(gx).Bind(&open)
	app.Response(open.GetList())
}

func (app Controller) DepositRecord(gx *gin.Context) {
	drp := services2.DepositRecordReq{}
	app.MakeContext(gx).Bind(&drp)
	app.Response(drp.GetList(auth.GetUserInfo(gx)))

}

func (app Controller) GameCategory(gx *gin.Context) {
	app.MakeContext(gx)
	app.Response(services2.GetGameCategory())
}

func (app Controller) GameList(gx *gin.Context) {
	s := services2.GameList{}
	app.MakeContext(gx).Bind(&s)
	app.Response(s.Func())
}

func (app Controller) Withdraw(gx *gin.Context) {
	withdraw := &services2.Withdraw{}
	app.MakeContext(gx).Bind(&withdraw).GetUserInfo()
	app.Response(withdraw.WithdrawFunc(app.UserInfo))
}
func (app Controller) Bet(gx *gin.Context) {
	bet := &services2.Bet{}
	gx.BindJSON(bet)
	app.MakeContext(gx).Bind(&bet).GetUserInfo()
	app.Response(bet.BetFunc(app.UserInfo))

}
