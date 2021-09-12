package showuserwalletinfo

import (
  "context"
  "mywallet/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showUserWalletInfoInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowUserWalletInfo
func NewUsecase(outputPort Outport) Inport {
  return &showUserWalletInfoInteractor{
    outport: outputPort,
  }
}

// Execute the usecase ShowUserWalletInfo
func (r *showUserWalletInfoInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  err := repository.WithoutTrx(ctx, r.outport, func(ctx context.Context) error {

    walletObjs, err := r.outport.FindAllWalletByUser(ctx, req.UserID)
    if err != nil {
      return err
    }

    res.Wallets = walletObjs

    csh, err := r.outport.FindAllCardSpendHistory(ctx, req.UserID)
    if err != nil {
      return err
    }

    res.CardSpendHistories = csh

    return nil
  })

  if err != nil {
    return nil, err
  }

  return res, nil
}
