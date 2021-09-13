package addnewcard

import (
  "context"
)

// Inport of AddnewCard
type Inport interface {
  Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase AddnewCard
type InportRequest struct {
  UserID        string
  WalletID      string
  CardName      string
  LimitAmount   float64
  LimitDuration string
}

// InportResponse is response payload after running the usecase AddnewCard
type InportResponse struct {
  CardID string
}
