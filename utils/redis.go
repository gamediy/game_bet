package utils

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var RedisMain *redis.Client

func redisInit() {
	RedisMain = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       0,                                 // use default DB
	})

}
func RedisGet[T any](key string, data *T) {
	result, err := RedisMain.Get(context.Background(), key).Result()
	if err != nil {
		return
	}
	json.Unmarshal([]byte(result), data)

}
func RedisSet[T any](key string, data T, exp time.Duration) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = RedisMain.Set(context.Background(), key, marshal, exp).Err()
	if err != nil {
		return err
	}
	return nil

}
