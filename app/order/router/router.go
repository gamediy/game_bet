package router

import (
	"bet/app/order/controller"
	"bet/core/auth"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	engine.Use(auth.Cors())
	api := engine.Group("/api/order")
	api.Use(auth.GinJWTMiddleware().MiddlewareFunc())
	{
		api.POST("/withdraw", controller.Withdraw)
		api.POST("/deposit", controller.Deposit)
		api.POST("/bet", controller.Bet)

	}
	engine.Run(":8082")
}
