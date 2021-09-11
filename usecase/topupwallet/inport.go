package topupwallet

import (
  "context"
  "time"
)

// Inport of TopupWallet
type Inport interface {
  Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase TopupWallet
type InportRequest struct {
  UserID   string
  WalletID string
  Amount   float64
  Date     time.Time
}

// InportResponse is response payload after running the usecase TopupWallet
type InportResponse struct {
}
