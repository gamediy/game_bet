package db

import (
	"bet/config"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var RedisMain *redis.Client

func init() {
	config.ConfigInit()
	RedisMain = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       0,                                 // use default DB
	})

}
func RedisGet[T any](key string, data *T) error {
	result, err := RedisMain.Get(context.Background(), key).Result()
	if err != nil {

		return err
	}
	if result == "" {
		return errors.New("no data")
	}
	err = json.Unmarshal([]byte(result), data)
	if err != nil {
		return err
	}
	return nil

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
