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
	api.Use(auth.GinJWTMiddleware().MiddlewareFunc())
	{
		api.POST("/amount_list", controller.AmountList)
		api.GET("/issue", controller.Issue)
		api.POST("/game_open", controller.GameOpen)
		api.POST("/deposit_record", controller.DepositRecord)
		api.POST("/game_category", controller.GameCategory)

	}
	engine.Run(":8083")
}
