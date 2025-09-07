package httptransport

import (
	"context"
	"encoding/json"
	"net/http"
	"tx-api/config"
	"tx-api/errors"

	"tx-api/core/kit"
	"tx-api/endpoint"

	gkhttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func decodeCreateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.NewErrInvalidJsonRequest()
	}
	return request, nil

}

func NewHTTPHandler(endpoints endpoint.Set) http.Handler {
	options := []gkhttp.ServerOption{
		gkhttp.ServerErrorEncoder(kit.ErrorEncoder),
		gkhttp.ServerBefore(kit.ExtractAuthTokenContext(&config.Config{})),
	}

	r := mux.NewRouter()

	accountCreateHandler := gkhttp.NewServer(
		endpoints.CreateAccountEndpoint,
		decodeCreateAccountRequest,
		kit.GenericEncodeResponse,
		options...,
	)
	r.Handle("/accounts", accountCreateHandler).Methods("POST")

	return r
}
