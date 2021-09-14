package showuserwalletinfo

import (
	"context"
	"mywallet/domain/entity"
)

// Inport of ShowAllUserWallet
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAllUserWallet
type InportRequest struct {
	UserID string
}

// InportResponse is response payload after running the usecase ShowAllUserWallet
type InportResponse struct {
  Wallets            []*entity.Wallet
  CardSpendHistories map[string]*entity.CardSpendHistory
}
