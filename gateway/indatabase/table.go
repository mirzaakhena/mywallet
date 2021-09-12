package indatabase

import (
  "mywallet/domain/vo"
  "time"
)

type UserTable struct {
  ID string `gorm:"primaryKey"` //
}

type CardTable struct {
  ID            string  `gorm:"primaryKey"`
  LimitAmount   float64 ``
  LimitDuration string  `gorm:"size:8"`
}

type WalletTable struct {
  ID          string `` //
  User        *UserTable
  UserTableID string
  Balance     vo.Money
}

type WalletCardTable struct {
  WalletTableID string
  Wallet        *WalletTable
}

type CardSpendHistoryTable struct {
  TransactionID    string `gorm:"primaryKey"`
  User             *UserTable
  UserTableID      string
  Card             *CardTable
  CardTableID      string
  Amount           float64 // currently it is not very useful here
  BalanceRemaining float64
  Date             time.Time
}
