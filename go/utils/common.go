package utils

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"math"
	"net/http"
)

const Md5Key = "u8r7XR1z"

func Md5(data string) string {
	// 进行md5加密，因为Sum函数接受的是字节数组，因此需要注意类型转换
	srcCode := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", srcCode)
}

func SnowflakeId() int64 {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	// Generate a snowflake ID.
	id := node.Generate().Int64()
	return id
}

func Md5Encryption(text string) string {
	return Md5(Md5Key + text)
}

// 将float64转成精确的int64
func Float64ToInt64(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

// 将int64恢复成正常的float64
func Int64ToFloat64(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

// 精准float64
func PreciseToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

// 精准int64
func PreciseToInt64(num int64, retain int) int64 {
	return int64(Int64ToFloat64(num, retain))
}

func NewRateLimiter(path string, useSubject bool, limit redis_rate.Limit) func(c *gin.Context) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_ = rdb.FlushDB(ctx).Err()
	return func(c *gin.Context) {

		key := path
		if useSubject {
			key += ":" + c.Request.Header.Get("Authorization")
		}
		ctx := context.Background()
		limiter := redis_rate.NewLimiter(rdb)
		res, err := limiter.Allow(ctx, key, limit)
		if err != nil {
			panic(err)
		}

		log.Debug("allowed: ", res.Allowed, ", remaining: ", res.Remaining)
		if res.Allowed == 0 {
			r := Result[string]{
				Code:    500,
				Message: "limit",
			}
			c.JSON(http.StatusOK, r)
			c.Abort()
			return
		}
	}
}
