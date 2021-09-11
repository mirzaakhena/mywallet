package entity

import (
	"mywallet/domain/vo"
	"time"
)

type CardSpendHistory struct {
	TransactionID    string
	User             *User
	Card             *Card
	Amount           vo.Money // currently it is not very useful here
	BalanceRemaining vo.Money
	Date             time.Time
}


