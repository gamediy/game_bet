package main

import (
	"bet/app/settle/services"
	"bet/utils"
	"fmt"
)

func main() {
	utils.Init()

	settle := services.Settle{
		GameName:   "",
		GameCode:   1000,
		OpenResult: 1,
	}
	settle.Calc()
	///init_play.PlayList[play.TwoPalyType].Settle()
	fmt.Println("settle")
}
