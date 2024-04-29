package main

import (
	"bet/app/game/router"
	"bet/db"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println(viper.Get("env"))
	defer db.SqlDB.Close()
	router.Run()

}
