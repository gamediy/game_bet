package router

import (
	"bet/core/auth"
	"bet/db"
	"bet/utils"
	"fmt"
	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-gonic/gin"
	"time"
)

func RateLimiter() gin.HandlerFunc {
	return ratelimiter.GinRedisRatelimiter(db.RedisMain, ratelimiter.GinRatelimiterConfig{
		// config: how to generate a limit key
		LimitKey: func(c *gin.Context) string {
			uid := c.Keys["Uid"]
			info := uid.(*auth.UserInfo)
			key := fmt.Sprintf("RateLimiter:%v_%v_%v", info.Uid)
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
			return time.Microsecond * 1, 3
		},
	})
}
