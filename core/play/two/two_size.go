package two

import (
	"bet/core/play"
	"errors"
	"strconv"
)

type TwoSize struct {
}

func (TwoSize) Won(openResult interface{}, betContent string, won *play.Won) {
	w := play.WonItem{
		PlayCode: play.SmallPlayCode,
		PlayName: "Small",
	}
	open := openResult.(int)
	if open > 4 {
		w.PlayCode = play.BigPlayCode
		w.PlayName = "Big"
	}

	won.List = append(won.List, w)

}
func (TwoSize) Check(bet string) error {
	i, err := strconv.ParseUint(bet, 10, 64)
	if err != nil {
		return err
	}
	if i > 9 {
		return errors.New("Wrong bet content")
	}
	return nil
}
