package addnewwallet

import (
  "context"
  "mywallet/application/apperror"
  "mywallet/domain/entity"
)

//go:generate mockery --name Outport -output mocks/

type addNewWalletInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase AddNewWallet
func NewUsecase(outputPort Outport) Inport {
  return &addNewWalletInteractor{
    outport: outputPort,
  }
}

// Execute the usecase AddNewWallet
func (r *addNewWalletInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  cardObj, err := entity.NewCard(entity.CardRequest{
    CardName:      req.CardName,
    LimitAmount:   req.LimitAmount,
    LimitDuration: req.LimitDuration,
  })
  if err != nil {
    return nil, err
  }

  userObj, err := r.outport.FindUserByID(ctx, req.UserID)
  if err != nil {
    return nil, err
  }
  if userObj == nil {
    return nil, apperror.UserIsNotFound
  }

  walletObj, err := entity.NewWallet(entity.WalletRequest{
    WalletName: req.WalletName,
    User:       userObj,
    Card:       cardObj,
  })
  if err != nil {
    return nil, err
  }

  err = r.outport.SaveWallet(ctx, walletObj)
  if err != nil {
    return nil, err
  }

  return res, nil
}