package model

type UserMoney struct {
	Uid           int64 `gorm:"primarykey"`
	Balance       float32
	TotalBet      float32
	TotalDeposit  float32
	TotalWithdraw int
	Account       string
	Freeze        float32
	Email         string
}

func (UserMoney) TableName() string {
	return "user_money"
}
