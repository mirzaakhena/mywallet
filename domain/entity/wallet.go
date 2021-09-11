package entity

import (
  "mywallet/application/apperror"
  "mywallet/domain/vo"
  "strings"
  "time"
)

type Wallet struct {
  ID      string `` //
  User    *User
  Balance vo.Money
  Cards   []*Card
}

type WalletRequest struct {
  WalletName string
  User       *User
  Card       *Card
}

func NewWallet(req WalletRequest) (*Wallet, error) {

  if req.WalletName == "" {
    return nil, apperror.WalletNameMustNotEmpty
  }

  if req.User == nil {
    return nil, apperror.UserMustNotNil
  }

  if req.Card == nil {
    return nil, apperror.CardMustNotNil
  }

  var obj Wallet
  obj.ID = strings.ToLower(strings.ReplaceAll(req.WalletName, " ", ""))
  obj.User = req.User
  obj.Cards = append(obj.Cards, req.Card)
  obj.Balance = vo.NewMoneyZero()

  return &obj, nil
}

func (w *Wallet) AddCard(card *Card) error {

  for _, c := range w.Cards {
    if c.ID == card.ID {
      return apperror.CardIDIsExist.Var(card.ID)
    }
  }

  w.Cards = append(w.Cards, card)
  return nil
}

func (w *Wallet) Topup(amount vo.Money) error {
  w.Balance += amount
  return nil
}

func (w *Wallet) Spend(lastCardSpendHistory *CardSpendHistory, amount vo.Money, cardID string, now time.Time) (*CardSpendHistory, error) {

  card := w.findCard(cardID)

  // user never used the balance yet
  if lastCardSpendHistory == nil {
    return w.newLimitToSpend(amount, card, now)
  }

  // user already use the balance, so we need to check the limit time
  stillPossibleToSpend, err := w.stillPossibleToSpend(lastCardSpendHistory.Date, now, card.LimitDuration)
  if err != nil {
    return nil, err
  }

  if stillPossibleToSpend {
    return w.spendRemainingBalance(lastCardSpendHistory.BalanceRemaining, amount, card, now)
  }

  // it is in the next time
  return w.newLimitToSpend(amount, card, now)

}

func (w *Wallet) spendRemainingBalance(remainingBalance, amountToSpend vo.Money, card *Card, now time.Time) (*CardSpendHistory, error) {

  if remainingBalance == 0 {
    return nil, apperror.CardLimitReachZero
  }

  if amountToSpend > remainingBalance {
    return nil, apperror.AmountGreaterThanRemainingBalanceInCard
  }

  newCardSpend := CardSpendHistory{
    User:             w.User,
    Card:             card,
    Amount:           amountToSpend,
    BalanceRemaining: remainingBalance - amountToSpend,
    Date:             now,
  }

  w.Balance = w.Balance - amountToSpend

  return &newCardSpend, nil
}

func (w *Wallet) stillPossibleToSpend(lastDate, now time.Time, limitTime vo.LimitTime) (bool, error) {

  if limitTime == vo.DailyLimitTimeEnum {
    stillPossibleToSpend, err := w.inTheSameDay(lastDate, now)
    if err != nil {
      return false, err
    }

    return stillPossibleToSpend, nil
  }


  if limitTime == vo.WeeklyLimitTimeEnum {
    stillPossibleToSpend, err := w.inTheSameWeek(lastDate, now)
    if err != nil {
      return false, err
    }
    return stillPossibleToSpend, nil

  }

  if limitTime == vo.MonthlyLimitTimeEnum {
    stillPossibleToSpend, err := w.inTheSameMonth(lastDate, now)
    if err != nil {
      return false, err
    }

    return stillPossibleToSpend, nil
  }

  return false, apperror.UnrecognizedLimitTime

}

func (w *Wallet) findCard(cardID string) *Card {
  for _, c := range w.Cards {
    if c.ID == cardID {
      return c
    }
  }
  return nil
}

func (w *Wallet) newLimitToSpend(amount vo.Money, card *Card, now time.Time) (*CardSpendHistory, error) {

  if amount > w.Balance {
    return nil, apperror.AmountGreaterThanBalance
  }

  if amount > card.LimitAmount {
    return nil, apperror.AmountGreaterThanLimitInCard
  }

  newUserSpend := CardSpendHistory{
    User:             w.User,
    Card:             card,
    Amount:           amount,
    BalanceRemaining: card.LimitAmount - amount,
    Date:             now,
  }

  w.Balance = w.Balance - amount

  return &newUserSpend, nil

}

func (w *Wallet) inTheSameDay(lastDate time.Time, now time.Time) (bool, error) {

  err := w.validateNowAsFuture(lastDate, now)
  if err != nil {
    return false, err
  }

  y1, m1, d1 := lastDate.Date()
  y2, m2, d2 := now.Date()

  return y1 == y2 && m1 == m2 && d1 == d2, nil
}

func (w *Wallet) inTheSameWeek(lastDate time.Time, now time.Time) (bool, error) {

  err := w.validateNowAsFuture(lastDate, now)
  if err != nil {
    return false, err
  }

  y1, w1 := lastDate.ISOWeek()
  y2, w2 := lastDate.ISOWeek()

  return y1 == y2 && w1 == w2, nil
}

func (w *Wallet) inTheSameMonth(lastDate time.Time, now time.Time) (bool, error) {

  err := w.validateNowAsFuture(lastDate, now)
  if err != nil {
    return false, err
  }

  y1, m1, _ := lastDate.Date()
  y2, m2, _ := now.Date()

  return y1 == y2 && m1 == m2, nil
}

func (w *Wallet) validateNowAsFuture(lastDate time.Time, now time.Time) error {
  if now.Before(lastDate) {
    return apperror.DateNowMustFutureFromLastDate
  }
  return nil
}