package main

import (
	"bet/app/user/router"
	"fmt"
	"github.com/spf13/viper"
)

var envArgs string

func main() {

	fmt.Println(viper.Get("env"))

	router.Run()

}
