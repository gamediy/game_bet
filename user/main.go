package main

import (
	"bet/user/router"
	"bet/utils"
	"fmt"
	"github.com/spf13/viper"
)

var envArgs string

func main() {

	fmt.Println(viper.Get("env"))
	utils.Init()

	router.Run()

}
