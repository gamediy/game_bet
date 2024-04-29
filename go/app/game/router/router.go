package router

import (
	"bet/app/game/controller"
	"bet/app/router"
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

		api.Use(router.RateLimiter()).POST("/amount_list", c.AmountList)
		api.Use(router.RateLimiter()).POST("/issue", c.Issue)
		api.Use(router.RateLimiter()).POST("/game_open", c.GameOpen)
		api.Use(router.RateLimiter()).POST("/deposit_record", c.DepositRecord)
		api.Use(router.RateLimiter()).POST("/game_category", c.GameCategory)
		api.Use(router.RateLimiter()).POST("/game_list", c.GameList)
		api.Use(router.RateLimiter()).POST("/withdraw", c.Withdraw)
		api.Use(router.RateLimiter()).POST("/bet", c.Bet)

	}
	engine.Run(":8083")
}
