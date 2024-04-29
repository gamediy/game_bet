package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var once sync.Once

func ConfigInit() {
	once.Do(func() {
		viper.AutomaticEnv()
		env := viper.Get("env")
		if env == nil {
			env = "dev"
			fmt.Println("please set env")
		}
		viper.SetConfigName(env.(string))
		viper.SetConfigType("toml")
		viper.AddConfigPath("./config") //搜索路径可以设置多个，viper 会根据设置顺序依次查找
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("read config failed: %v", err)
		}
	})

}
