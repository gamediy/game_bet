package controller

import (
	"bet/app/controller"
	"bet/app/user/services"
	"bet/core/auth"
	"bet/model"
	"bet/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	controller.Controller
}

func (app *Controller) Register(gx *gin.Context) {
	var register = services.Register{}
	err := app.MakeContext(gx).Bind(&register).Errors
	if err != nil {
		app.Error("", err.Error())
		return
	}

	app.Response(register.Logic())

}

func UserInfo(gx *gin.Context) {

	uid := gx.Keys["Uid"]
	info := uid.(*auth.UserInfo)
	user := model.UserAmount{}
	utils.DB.Table("user_amount").First(&user, info.Uid)
	gx.JSON(http.StatusOK, gin.H{
		"uid":     info.Uid,
		"account": info.Account,
		"balance": user.Balance,
	})

}
