package utils

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var DB *gorm.DB

var err error

var once sync.Once

func init() {
	viper.AutomaticEnv()

}
func mysqlInit() {
	DB, err = gorm.Open(mysql.Open(viper.GetString("database.main")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func Init() {
	once.Do(func() {

		configInit()
		mysqlInit()
		redisInit()
		InitTrans("en")

	})
}

func configInit() {
	env := viper.Get("env")
	viper.SetConfigName(env.(string))
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config") //搜索路径可以设置多个，viper 会根据设置顺序依次查找
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err)
	}

}
