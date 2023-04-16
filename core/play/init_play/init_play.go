package init_play

import (
	"bet/core/play"
	"bet/core/play/two"
)

var PlayList map[int32]play.Play

func init() {
	PlayList = make(map[int32]play.Play)
	t := two.TwoSize{}
	PlayList[play.TwoSize] = &t

}
