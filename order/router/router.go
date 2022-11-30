package router

import (
	"bet/core"
	"bet/order/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	engine.Use(core.Cors())
	auth := engine.Group("/api/order")
	auth.Use(core.GinJWTMiddleware().MiddlewareFunc())
	{
		auth.POST("/withdraw", controller.Withdraw)
		auth.POST("/deposit", controller.Deposit)
		auth.GET("/list", controller.AmountList)

	}
	engine.Run(":8082")
}
