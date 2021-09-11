package showalluser

import "context"

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

	userObjs, err := r.outport.FindAllUser(ctx)
	if err != nil {
		return nil, err
	}

	res.Users = userObjs

	return res, nil
}
