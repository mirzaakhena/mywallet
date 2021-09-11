package vo

import "mywallet/application/apperror"

type Money float64

func NewMoney(amount float64) (m Money, err error) {
  if amount < 0 {
    return 0, apperror.MoneyMustGreaterThanZero
  }
  m = Money(amount)
  return
}

func NewMoneyZero() Money {
  return Money(0)
}