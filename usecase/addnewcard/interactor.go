package addnewcard

import (
  "context"
  "mywallet/application/apperror"
  "mywallet/domain/entity"
)

//go:generate mockery --name Outport -output mocks/

type addnewCardInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase AddnewCard
func NewUsecase(outputPort Outport) Inport {
  return &addnewCardInteractor{
    outport: outputPort,
  }
}

// Execute the usecase AddnewCard
func (r *addnewCardInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  walletObj, err := r.outport.FindWalletByID(ctx, req.WalletID)
  if err != nil {
    return nil, err
  }

  if walletObj.User.ID != req.UserID {
    return nil, apperror.UserIDDoesNotMatch
  }

  cardObj, err := entity.NewCard(entity.CardRequest{
    CardName:      req.CardName,
    LimitAmount:   req.LimitAmount,
    LimitDuration: req.LimitDuration,
  })
  if err != nil {
    return nil, err
  }

  err = walletObj.AddCard(cardObj)
  if err != nil {
    return nil, err
  }

  err = r.outport.SaveWallet(ctx, walletObj)
  if err != nil {
    return nil, err
  }

  return res, nil
}
