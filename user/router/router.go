package router

import (
	"bet/core"
	"bet/user/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func Run() {
	engine := gin.Default()
	engine.Use(Cors())
	engine.POST("/api/user/login", core.GinJWTMiddleware().LoginHandler)
	engine.POST("/api/user/register", controller.Register)

	auth := engine.Group("/api/user")
	auth.Use(core.GinJWTMiddleware().MiddlewareFunc())
	{
		auth.POST("/userinfo", controller.UserInfo)
	}
	engine.Run(":8081")
}
