package play

const (
	TwoSize_Big   = "BIG"
	TwoSize_Small = "SMALL"
)
const (
	TwoSize = "TwoSize"
)

type Play interface {
	Won(openResult interface{}, betContent interface{}, won *Won) //结算
	Check(betContent interface{}) error
}

type WonItem struct {
	PlayCode string
}
type Won struct {
	GameCode int32
	GameName string

	List []WonItem
}
