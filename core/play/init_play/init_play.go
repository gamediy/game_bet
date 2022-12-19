package init_play

import (
	"bet/core/play"
	"bet/core/play/two"
)

var PlayList map[string]play.Play

func init() {
	PlayList = make(map[string]play.Play)
	t := two.TwoSize{}
	PlayList[play.TwoPalyType] = &t

}
