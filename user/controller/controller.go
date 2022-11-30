package controller

import (
	"bet/core"
	"bet/user/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(gx *gin.Context) {
	var register = &services.Register{}
	gx.BindJSON(register)
	gx.JSON(http.StatusOK, register.Logic())

}

func UserInfo(gx *gin.Context) {

	uid := gx.Keys["Uid"]
	info := uid.(*core.UserInfo)
	gx.JSON(http.StatusOK, gin.H{
		"uid":     info.Uid,
		"account": info.Account,
	})

}
