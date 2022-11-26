package router

import (
	"bet/core"
	"bet/user/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()

	engine.POST("/api/user/login", core.GinJWTMiddleware().LoginHandler)
	engine.POST("/api/user/register", controller.Register)

	auth := engine.Group("/api/user")
	auth.Use(core.GinJWTMiddleware().MiddlewareFunc())
	{
		auth.POST("/userinfo", controller.UserInfo)
	}
	engine.Run(":8081")
}
