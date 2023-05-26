package main

import (
	"bet/app/user/router"
	"bet/db"
	"fmt"
	"github.com/spf13/viper"
)

var envArgs string

func main() {

	fmt.Println(viper.Get("env"))
	defer db.SqlDB.Close()
	router.Run()

}
