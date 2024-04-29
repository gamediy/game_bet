package two

import (
	"bet/core/play"
	"errors"
)

type TwoSize struct {
}

func (TwoSize) Won(openResult interface{}, betContent interface{}, won *play.Won) {
	w := play.WonItem{
		PlayCode: play.TwoSize_Small,
	}
	open := openResult.(int)
	if open > 4 {
		w.PlayCode = play.TwoSize_Big

	}

	won.List = append(won.List, w)

}
func (TwoSize) Check(bet interface{}) error {
	i := bet.(int)
	if i > 9 {
		return errors.New("Wrong bet content")
	}
	return nil
}
