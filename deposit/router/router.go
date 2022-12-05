package router

import (
	"bet/core/auth"
	"bet/deposit/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	engine.Use(auth.Cors())
	api := engine.Group("/api/deposit")
	api.Use(auth.GinJWTMiddleware().MiddlewareFunc())
	{

		api.POST("/deposit", controller.Deposit)

	}
	engine.Run(":8084")
}
