package play

const (
	BigPlayCode   = 100
	SmallPlayCode = 101
)
const (
	TwoPalyType = "Two"
)

type Play interface {
	Won(openResult interface{}, betContent string, won *Won) //结算
	Check(betContent string) error
}

type WonItem struct {
	PlayName string
	PlayCode int32
}
type Won struct {
	GameCode int32
	GameName string

	List []WonItem
}
