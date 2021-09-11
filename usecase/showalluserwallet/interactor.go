package showalluserwallet

import "context"

//go:generate mockery --name Outport -output mocks/

type showAllUserWalletInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAllUserWallet
func NewUsecase(outputPort Outport) Inport {
	return &showAllUserWalletInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ShowAllUserWallet
func (r *showAllUserWalletInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	walletObjs, err := r.outport.FindAllWalletByUser(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	res.Wallets = walletObjs

	return res, nil
}
