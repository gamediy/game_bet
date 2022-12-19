package router

import (
	"bet/core/auth"
	"bet/game/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	engine.Use(auth.Cors())
	api := engine.Group("/api/game")
	api.Use(auth.GinJWTMiddleware().MiddlewareFunc())
	{
		api.POST("/amount_list", controller.AmountList)
		api.GET("/issue", controller.Issue)
		api.POST("/game_open", controller.GameOpen)
		api.POST("/deposit_record", controller.DepositRecord)
		api.POST("/game_list", controller.GameList)
	}
	engine.Run(":8083")
}
