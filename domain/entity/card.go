package entity

import (
  "mywallet/application/apperror"
  "mywallet/domain/vo"
  "strings"
)

type Card struct {
  ID            string `` //
  LimitAmount   vo.Money
  LimitDuration vo.LimitTime
}

type CardRequest struct {
  CardName      string
  LimitAmount   float64
  LimitDuration string
}

func NewCard(req CardRequest) (obj *Card, err error) {

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

  obj.ID = strings.ToLower(strings.ReplaceAll(req.CardName, " ", ""))
  obj.LimitAmount = limitAmount
  obj.LimitDuration, err = vo.NewLimitTime(req.LimitDuration)
  if err != nil {
    return nil, err
  }

  return obj, nil
}
