package order_bet_status

const (
	Bet  = 1
	Won  = 2
	Lost = 3
	Draw = 4
)

var statusMap = map[int32]string{
	1: "bet",
	2: "won",
	3: "lost",
	4: "draw",
}

func GetStatusStr(status int32) string {
	s, ok := statusMap[status]
	if ok {
		return s
	}
	return ""
}
