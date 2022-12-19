package main

import (
	"bet/game/router"
	"bet/utils"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println(viper.Get("env"))
	utils.Init()
	defer utils.SqlDb.Close()
	router.Run()

}
