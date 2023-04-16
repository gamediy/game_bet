package main

import (
	"bet/core/auth"

	"bet/socket/model"
	"bet/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	Upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	utils.Init()
	defer utils.SqlDb.Close()
	engine := gin.Default()
	engine.Use(auth.Cors())
	api := engine.Group("/api/socket")
	api.Use(auth.GinJWTMiddleware().MiddlewareFunc())
	{
		api.GET("/ws", func(ctx *gin.Context) {
			wsConn, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
			u := utils.Result[string]{
				Code:    500,
				Message: "",
			}
			if err != nil {
				u.Message = err.Error()
				ctx.JSON(http.StatusBadGateway, u)
			}
			model.NewUser(wsConn, auth.GetUserInfo(ctx))
			u.Code = 200
			u.Message = "success"
		})
	}
	engine.Run(":8085")
}
