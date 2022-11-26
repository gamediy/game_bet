package controller

import (
	"bet/core"
	"bet/user/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(gx *gin.Context) {
	var register = &services.Register{}
	gx.BindJSON(register)
	gx.JSON(http.StatusOK, register.Logic())

}
func Login(gx *gin.Context) {
	var login = &services.Login{}
	gx.BindJSON(login)
	gx.JSON(http.StatusOK, login.Login)
}
func UserInfo(gx *gin.Context) {
	fmt.Println("11")
	uid := gx.Keys["Uid"]
	info := uid.(*core.UserInfo)
	gx.JSON(http.StatusOK, gin.H{
		"uid":     info.Uid,
		"account": info.Account,
	})
	fmt.Println(gx)

}
