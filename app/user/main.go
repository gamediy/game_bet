package main

import (
	"bet/app/user/router"
	"bet/utils"
	"fmt"
	"github.com/spf13/viper"
)

var envArgs string

func main() {

	fmt.Println(viper.Get("env"))
	utils.Init()

	defer utils.SqlDb.Close()
	router.Run()

}
