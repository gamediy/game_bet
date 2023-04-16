package play

const (
	TwoSize_Big   = 1000
	TwoSize_Small = 1001
)
const (
	TwoSize = 100
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
