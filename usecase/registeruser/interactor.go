package registeruser

import (
  "context"
  "mywallet/domain/entity"
)

//go:generate mockery --name Outport -output mocks/

type registerUserInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase RegisterUser
func NewUsecase(outputPort Outport) Inport {
  return &registerUserInteractor{
    outport: outputPort,
  }
}

// Execute the usecase RegisterUser
func (r *registerUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  userObj, err := entity.NewUser(entity.UserRequest{Name: req.Name})
  if err != nil {
    return nil, err
  }

  err = r.outport.SaveUser(ctx, userObj)
  if err != nil {
    return nil, err
  }

  return res, nil
}
