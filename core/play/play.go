package play

const (
	BigPlayCode   = 100
	SmallPlayCode = 101
)
const (
	TwoPalyType = "Two"
)

type Play interface {
	Settle() []Won //结算

}

type Won struct {
	PlayName string
	PlayCode int32
}
