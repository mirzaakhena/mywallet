package entity

import (
  "mywallet/application/apperror"
  "mywallet/domain/vo"
  "time"
)

type Wallet struct {
  ID      string
  Name    string
  User    *User
  UserID  string
  Balance vo.Money
  Cards   []*Card
}

type WalletRequest struct {
  ID         string
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
  obj.ID = req.ID
  obj.Name = req.WalletName
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

//func (w *Wallet) Spend(lastCardSpendHistory *CardSpendHistory, amount vo.Money, cardID string, now time.Time) (*CardSpendHistory, error) {
//
//  card := w.FindCard(cardID)
//
//  // user never used the balance yet
//  if lastCardSpendHistory == nil {
//    return w.NewLimitToSpend(amount, card, now)
//  }
//
//  // user already use the balance, so we need to check the limit time
//  stillPossibleToSpend, err := lastCardSpendHistory.IsStillPossibleToSpend(now)
//  if err != nil {
//    return nil, err
//  }
//
//  if stillPossibleToSpend {
//    return w.SpendRemainingBalance(lastCardSpendHistory.BalanceRemaining, amount, card, now)
//  }
//
//  // it is in the next time
//  return w.NewLimitToSpend(amount, card, now)
//
//}

func (w *Wallet) SpendRemainingBalance(remainingBalance, amountToSpend vo.Money, card *Card, now time.Time) (*CardSpendHistory, error) {

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

func (w *Wallet) FindCard(cardID string) (*Card, error) {
  for _, c := range w.Cards {
    if c.ID == cardID {
      return c, nil
    }
  }
  return nil, apperror.CardNotFound
}

func (w *Wallet) NewLimitToSpend(amount vo.Money, card *Card, now time.Time) (*CardSpendHistory, error) {

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
