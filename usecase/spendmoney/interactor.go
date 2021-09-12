package spendmoney

import (
  "context"
  "mywallet/application/apperror"
  "mywallet/domain/repository"
  "mywallet/domain/vo"
)

//go:generate mockery --name Outport -output mocks/

type spendMoneyInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase SpendMoney
func NewUsecase(outputPort Outport) Inport {
  return &spendMoneyInteractor{
    outport: outputPort,
  }
}

// Execute the usecase SpendMoney
func (r *spendMoneyInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  err := repository.WithTrx(ctx, r.outport, func(ctx context.Context) error {

    walletObj, err := r.outport.FindWalletByID(ctx, req.WalletID)
    if err != nil {
      return err
    }

    if walletObj.User.ID != req.UserID {
      return apperror.UserIDDoesNotMatch
    }

    amountToSpend, err := vo.NewMoney(req.Amount)
    if err != nil {
      return err
    }

    cardSpendHistoryObj, err := r.outport.FindLastCardSpendHistory(ctx, req.CardID)
    if err != nil {
      return err
    }

    ushObj, err := walletObj.Spend(cardSpendHistoryObj, amountToSpend, req.CardID, req.Date)
    if err != nil {
      return err
    }

    err = r.outport.UpdateWalletBalance(ctx, walletObj)
    if err != nil {
      return err
    }

    err = r.outport.SaveCardSpendHistory(ctx, ushObj)
    if err != nil {
      return err
    }

    return nil

  })

  if err != nil {
    return nil, err
  }

  return res, nil
}
