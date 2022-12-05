package two_size

import (
	"bet/core/play"
)

type Two struct {
	OpenResult int32
	BetContent int32
}

func (this *Two) Settle() []play.Won {
	won := play.Won{
		PlayCode: play.SmallPlayCode,
		PlayName: "Small",
	}
	wons := make([]play.Won, 0)
	wons = append(wons, won)
	if this.OpenResult > 4 {
		won.PlayCode = play.BigPlayCode
		won.PlayName = "Big"
		wons = append(wons, won)
	}

	return wons
}
