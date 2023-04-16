package router

import (
	"bet/app/deposit/controller"
	services2 "bet/app/deposit/services"
	"bet/core/auth"
	"bet/utils"
	"fmt"
	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

func Run() {
	engine := gin.Default()
	engine.Use(auth.Cors())
	api := engine.Group("/api/deposit")
	api.Use(auth.GinJWTMiddleware().MiddlewareFunc())
	{

		api.Use(ratelimiter.GinRedisRatelimiter(utils.RedisMain, ratelimiter.GinRatelimiterConfig{
			// config: how to generate a limit key
			LimitKey: func(c *gin.Context) string {
				text := ""
				deposit := services2.Deposit{}
				c.ShouldBindBodyWith(&deposit, binding.JSON)

				fmt.Println(text)
				uid := c.Keys["Uid"]
				info := uid.(*auth.UserInfo)
				key := fmt.Sprintf("RateLimiter:%v_%v_%v", info.Uid, deposit.AmountItemCode, deposit.Amount)
				return key
			},
			// config: how to respond when limiting
			LimitedHandler: func(c *gin.Context) {
				u := utils.Result[string]{
					Code:    500,
					Message: "too many requests!!!",
				}
				c.JSON(200, u)
				c.Abort()
				return
			},
			// config: return ratelimiter token fill interval and bucket size
			TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
				return time.Minute * 1, 1
			},
		})).POST("/deposit", controller.Deposit)

	}
	engine.Run(":8084")
}
