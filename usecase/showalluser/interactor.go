package showalluser

import (
  "context"
  "mywallet/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type showAllUserInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ShowAllUser
func NewUsecase(outputPort Outport) Inport {
  return &showAllUserInteractor{
    outport: outputPort,
  }
}

// Execute the usecase ShowAllUser
func (r *showAllUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  err := repository.WithoutTrx(ctx, r.outport, func(ctx context.Context) error {

    userObjs, err := r.outport.FindAllUser(ctx)
    if err != nil {
      return err
    }

    res.Users = userObjs

    return nil
  })

  if err != nil {
    return nil, err
  }

  return res, nil
}
