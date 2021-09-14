package showuserwalletinfo

import (
	"context"
	"mywallet/domain/entity"
	"mywallet/domain/repository"
	"time"
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

		res.CardSpendHistories = r.convertToMap(walletObjs, csh)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *showUserWalletInfoInteractor) convertToMap(wallets []*entity.Wallet, csh []*entity.CardSpendHistory) map[string]*entity.CardSpendHistory {

	mapCard := map[string]*entity.CardSpendHistory{}
	for _, c := range csh {
		mapCard[c.CardID] = c
	}

	for _, w := range wallets {
		for _, c := range w.Cards {
			_, exist := mapCard[c.ID]
			if !exist {
				mapCard[c.ID] = &entity.CardSpendHistory{
					ID:               "",
					User:             nil,
					UserID:           w.UserID,
					Card:             nil,
					CardID:           c.ID,
					Amount:           0,
					BalanceRemaining: c.LimitAmount,
					Date:             time.Time{},
				}
			}
		}
	}

	return mapCard
}
