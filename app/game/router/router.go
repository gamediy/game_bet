package router

import (
	"bet/app/game/controller"
	"bet/core/auth"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	engine.Use(auth.Cors())
	api := engine.Group("/api/game")
	c := controller.Controller{}
	api.Use(auth.GinJWTMiddleware().MiddlewareFunc())
	{
		api.POST("/amount_list", c.AmountList)
		api.POST("/issue", c.Issue)
		api.POST("/game_open", c.GameOpen)
		api.POST("/deposit_record", c.DepositRecord)
		api.POST("/game_category", c.GameCategory)
		api.POST("/game_list", c.GameList)
		api.POST("/withdraw", controller.Withdraw)
		api.POST("/deposit", controller.Deposit)
		api.POST("/bet", controller.Bet)

	}
	engine.Run(":8083")
}
