package init_play

import (
	"bet/core/play"
	"bet/core/play/two_size"
)

var PlayList map[string]play.Play

func init() {
	PlayList = make(map[string]play.Play)
	t := two_size.Two{}
	PlayList[play.TwoPalyType] = &t

}
