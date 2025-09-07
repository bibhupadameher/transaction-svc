package kit

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"tx-api/core/apperror"
	"tx-api/errors"

	kitendpoint "github.com/go-kit/kit/endpoint"
)

func GenericEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitendpoint.Failer); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return errors.NewErrEncodingResponse()
	}
	return nil
}

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	httpStatus := http.StatusInternalServerError
	if xErr, ok := err.(apperror.AppError); ok {
		httpStatus = xErr.HTTPStatus
	} else if xErr, ok := err.(apperror.Errors); ok {
		if len(xErr) > 0 {
			httpStatus = xErr[0].HTTPStatus
		}

	}

	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(ErrorResponse{Error: err}); err != nil {
		log.Println(err)
	}
}

type ErrorResponse struct {
	Error error `json:"error"`
}
