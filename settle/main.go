package main

import (
	"bet/core/play"
	"bet/core/play/init_play"
	"fmt"
)

func main() {
	init_play.PlayList[play.TwoPalyType].Settle()

	fmt.Println("settle")
}
