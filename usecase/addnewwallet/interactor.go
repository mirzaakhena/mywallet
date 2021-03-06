package addnewwallet

import (
	"context"
	"mywallet/application/apperror"
	"mywallet/domain/entity"
	"mywallet/domain/repository"
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

	err := repository.WithTrx(ctx, r.outport, func(ctx context.Context) error {

		cardObj, err := entity.NewCard(entity.CardRequest{
			ID:            r.outport.GenerateID(ctx),
			CardName:      req.CardName,
			LimitAmount:   req.LimitAmount,
			LimitDuration: req.LimitDuration,
		})
		if err != nil {
			return err
		}

		userObj, err := r.outport.FindUserByID(ctx, req.UserID)
		if err != nil {
			return err
		}
		if userObj == nil {
			return apperror.UserIsNotFound
		}

		walletObj, err := entity.NewWallet(entity.WalletRequest{
			ID:         r.outport.GenerateID(ctx),
			WalletName: req.WalletName,
			User:       userObj,
			Card:       cardObj,
		})
		if err != nil {
			return err
		}

		err = r.outport.SaveWallet(ctx, walletObj)
		if err != nil {
			return err
		}

		err = r.outport.SaveCard(ctx, walletObj.ID, cardObj)
		if err != nil {
			return err
		}

		res.WalletID = walletObj.ID
		res.CardID = cardObj.ID

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
