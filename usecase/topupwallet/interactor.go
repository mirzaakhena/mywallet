package topupwallet

import (
	"context"
	"mywallet/application/apperror"
	"mywallet/domain/vo"
)

//go:generate mockery --name Outport -output mocks/

type topupWalletInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase TopupWallet
func NewUsecase(outputPort Outport) Inport {
	return &topupWalletInteractor{
		outport: outputPort,
	}
}

// Execute the usecase TopupWallet
func (r *topupWalletInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	walletObj, err := r.outport.FindWalletByID(ctx, req.WalletID)
	if err != nil {
		return nil, err
	}

	if walletObj.User.ID != req.UserID {
		return nil, apperror.UserIDDoesNotMatch
	}

	amountToTopup, err := vo.NewMoney(req.Amount)
	if err != nil {
		return nil, err
	}

	err = walletObj.Topup(amountToTopup)
	if err != nil {
		return nil, err
	}

	err = r.outport.UpdateWalletBalance(ctx, walletObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}
