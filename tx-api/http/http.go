package httptransport

import (
	"context"
	"encoding/json"
	"net/http"
	"tx-api/dto"
	"tx-api/errors"

	"tx-api/core/kit"
	"tx-api/endpoint"

	gkhttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func decodeCreateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.NewErrInvalidJsonRequest()
	}
	return request, nil

}

func decodeCreateTransactionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.NewErrInvalidJsonRequest()
	}
	return request, nil

}

func decodeGetAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	var accountID uuid.UUID
	id, ok := vars["accountId"]
	if ok {
		var err error
		accountID, err = uuid.Parse(id)
		if err != nil {
			return nil, errors.NewErrFieldMissingorInvalid("accountID")
		}
	} else {
		return nil, errors.NewErrFieldMissingorInvalid("accountID")
	}
	return endpoint.GetAccountRequest{GetAccountRequest: dto.GetAccountRequest{AccountID: accountID}}, nil
}

func NewHTTPHandler(endpoints endpoint.Set) http.Handler {
	options := []gkhttp.ServerOption{
		gkhttp.ServerErrorEncoder(kit.ErrorEncoder),
	}

	r := mux.NewRouter()

	accountCreateHandler := gkhttp.NewServer(
		endpoints.CreateAccountEndpoint,
		decodeCreateAccountRequest,
		kit.GenericEncodeResponse,
		options...,
	)
	r.Handle("/accounts", accountCreateHandler).Methods("POST")

	getAccountHandler := gkhttp.NewServer(
		endpoints.GetAccountEndpoint,
		decodeGetAccountRequest,
		kit.GenericEncodeResponse,
		options...,
	)
	r.Handle("/accounts/{accountId}", getAccountHandler).Methods("GET")

	txCreateHandler := gkhttp.NewServer(
		endpoints.CreateTransactionEndpoint,
		decodeCreateTransactionRequest,
		kit.GenericEncodeResponse,
		options...,
	)
	r.Handle("/transactions", txCreateHandler).Methods("POST")

	return r
}
