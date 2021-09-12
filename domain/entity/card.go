package entity

import (
  "mywallet/application/apperror"
  "mywallet/domain/vo"
)

type Card struct {
  ID            string
  WalletID      string
  Name          string
  LimitAmount   vo.Money
  LimitDuration vo.LimitTime
}

type CardRequest struct {
  ID            string `` //
  CardName      string
  LimitAmount   float64
  LimitDuration string
}

func NewCard(req CardRequest) (obj *Card, err error) {

  if req.ID == "" {
    return nil, apperror.CardIDMustNotEmpty
  }

  if req.CardName == "" {
    return nil, apperror.CardUserNameMustNotEmpty
  }

  limitAmount, err := vo.NewMoney(req.LimitAmount)
  if err != nil {
    return nil, err
  }

  if limitAmount == 0 {
    return nil, apperror.LimitAmountMustNotZero
  }

  obj.ID = req.ID
  obj.Name = req.CardName
  obj.LimitAmount = limitAmount
  obj.LimitDuration, err = vo.NewLimitTime(req.LimitDuration)
  if err != nil {
    return nil, err
  }

  return obj, nil
}
