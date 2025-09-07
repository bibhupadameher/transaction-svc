package endpoint

import (
	"tx-api/service"

	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	CreateAccountEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc service.TransactionServiceInterface) Set {

	createAccountEndpoint := MakeCreateAccountEndpoint(svc)
	//	createAccountEndpoint = AuthMiddleware()(createAccountEndpoint)
	createAccountEndpoint = LoggingMiddleware()(createAccountEndpoint)

	return Set{
		CreateAccountEndpoint: createAccountEndpoint,
	}
}
