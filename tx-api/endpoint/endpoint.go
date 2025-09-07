package endpoint

import (
	"context"

	"tx-api/dto"
	"tx-api/service"

	"github.com/go-kit/kit/endpoint"
)

// Request/Response
type CreateAccountRequest struct {
	dto.CreateAccountRequest
}
type CreateAccountResponse struct {
	dto.CreateAccountResponse
	Error error `json:"-"`
}

var (
	_ endpoint.Failer = CreateAccountResponse{}
)

func (r CreateAccountResponse) Failed() error { return r.Error }

// Make endpoint
func MakeCreateAccountEndpoint(s service.TransactionServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		resp, err := s.CreateAccount(ctx, req.CreateAccountRequest)
		if err != nil {
			return nil, err
		}
		return CreateAccountResponse{resp, err}, nil
	}
}
