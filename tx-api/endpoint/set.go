package endpoint

import (
	"tx-api/service"

	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	CreateAccountEndpoint     endpoint.Endpoint
	GetAccountEndpoint        endpoint.Endpoint
	CreateTransactionEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc service.TransactionServiceInterface) Set {

	createAccountEndpoint := MakeCreateAccountEndpoint(svc)
	createAccountEndpoint = LoggingMiddleware()(createAccountEndpoint)

	getAccountEndpoint := MakeGetAccountEndpoint(svc)
	getAccountEndpoint = LoggingMiddleware()(getAccountEndpoint)

	createTransactionEndpoint := MakeCreateTransactionEndpoint(svc)
	createTransactionEndpoint = LoggingMiddleware()(createTransactionEndpoint)

	return Set{
		CreateAccountEndpoint:     createAccountEndpoint,
		GetAccountEndpoint:        getAccountEndpoint,
		CreateTransactionEndpoint: createTransactionEndpoint,
	}
}
