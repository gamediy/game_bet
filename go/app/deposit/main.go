package main

import (
	"bet/app/deposit/router"
	"bet/utils"
)

func main() {
	utils.Init()
	router.Run()

}
