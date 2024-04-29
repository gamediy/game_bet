package db

import (
	"bet/config"
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB
var SqlDB *sql.DB
var err error

func init() {
	config.ConfigInit()
	config := mysql.Open(viper.GetString("database.main"))
	GormDB, err = gorm.Open(config, &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	SqlDB, _ = GormDB.DB()
	fmt.Sprintf("数据库初始处成功：%s", config)
	SqlDB.SetMaxOpenConns(1000)
	defer func() {
		SqlDB.Close()
	}()
}
