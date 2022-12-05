package main

import (
	"bet/deposit/router"
	"bet/utils"
)

func main() {
	utils.Init()
	router.Run()

}
