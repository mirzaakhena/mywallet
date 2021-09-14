package indatabase

import (
	"time"
)

type User struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

type Wallet struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	User    *User
	UserID  string
	Balance float64
}

type Card struct {
	ID            string `gorm:"primaryKey"`
	WalletID      string
	Name          string
	LimitAmount   float64
	LimitDuration string `gorm:"size:8"`
}

type CardSpendHistory struct {
	ID               string `gorm:"primaryKey"`
	User             *User
	UserID           string
	Card             *Card
	CardID           string
	Amount           float64
	BalanceRemaining float64
	Date             time.Time
}
