package showalluser

import (
	"context"
	"mywallet/domain/entity"
)

// Inport of ShowAllUser
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase ShowAllUser
type InportRequest struct {
}

// InportResponse is response payload after running the usecase ShowAllUser
type InportResponse struct {
	Users []*entity.User
}
