package addnewwallet

import (
  "context"
)

// Inport of AddNewWallet
type Inport interface {
  Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase AddNewWallet
type InportRequest struct {
  UserID        string
  WalletName    string
  CardName      string
  LimitAmount   float64
  LimitDuration string
}

// InportResponse is response payload after running the usecase AddNewWallet
type InportResponse struct {
  WalletID string
  CardID   string
}
