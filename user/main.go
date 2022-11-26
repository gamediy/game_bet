package main

import (
	"bet/user/router"
	"bet/utils"
	"context"
	"fmt"
	"github.com/spf13/viper"
)

var envArgs string

func main() {

	fmt.Println(viper.Get("env"))
	utils.Init()
	err := utils.RedisMain.Set(context.Background(), "tst", 111, -1).Err()
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	router.Run()

}
