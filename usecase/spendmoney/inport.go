package spendmoney

import (
	"context"
	"time"
)

// Inport of SpendMoney
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase SpendMoney
type InportRequest struct {
	UserID   string
	WalletID string
	CardID   string
	Amount   float64
	Date     time.Time
}

// InportResponse is response payload after running the usecase SpendMoney
type InportResponse struct {
	CardSpendHistoryID string
}
