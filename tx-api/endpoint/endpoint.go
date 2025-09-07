package endpoint

import (
	"context"

	"tx-api/dto"
	"tx-api/service"

	"github.com/go-kit/kit/endpoint"
)

type CreateAccountRequest struct {
	dto.CreateAccountRequest
}
type CreateAccountResponse struct {
	dto.CreateAccountResponse
	Error error `json:"-"`
}

type GetAccountRequest struct {
	dto.GetAccountRequest
}
type GetAccountResponse struct {
	dto.GetAccountResponse
	Error error `json:"-"`
}

type CreateTransactionRequest struct {
	dto.CreateTransactionRequest
}
type CreateTransactionResponse struct {
	dto.CreateTransactionResponse
	Error error `json:"-"`
}

var (
	_ endpoint.Failer = CreateAccountResponse{}
	_ endpoint.Failer = GetAccountResponse{}
	_ endpoint.Failer = CreateTransactionResponse{}
)

func (r CreateAccountResponse) Failed() error     { return r.Error }
func (r GetAccountResponse) Failed() error        { return r.Error }
func (r CreateTransactionResponse) Failed() error { return r.Error }

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

func MakeGetAccountEndpoint(s service.TransactionServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAccountRequest)
		resp, err := s.GetAccount(ctx, req.GetAccountRequest)
		if err != nil {
			return nil, err
		}
		return GetAccountResponse{resp, err}, nil
	}
}

func MakeCreateTransactionEndpoint(s service.TransactionServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTransactionRequest)
		resp, err := s.CreateTransaction(ctx, req.CreateTransactionRequest)
		if err != nil {
			return nil, err
		}
		return CreateTransactionResponse{resp, err}, nil
	}
}
