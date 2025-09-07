package errors

import (
	"fmt"
	"net/http"
	"tx-api/core/apperror"
)

var (
	errFieldMissingCode    = "A010001"
	errFieldMissingType    = "InvalidRequest"
	errFieldMissingMessage = "field missing in the request"
	errFieldMissingDetails = func(fieldName string) string {
		return fmt.Sprintf(
			"%s field missing in the request", fieldName)
	}
)

type ErrFieldMissing struct {
	apperror.AppError
}

func NewErrFieldMissing(fieldName string) ErrFieldMissing {
	return ErrFieldMissing{apperror.New(
		errFieldMissingCode,
		errFieldMissingType,
		errFieldMissingMessage,
		http.StatusBadRequest,
		errFieldMissingDetails(fieldName),
	)}
}

var (
	errInvalidJsonRequestCode    = "A010002"
	errInvalidJsonRequestType    = "InvalidRequest"
	errInvalidJsonRequestMessage = "unable to unmarshal json request"
	errInvalidJsonRequestDetails = "cannot marshal json bytes to request; please check request model"
)

type ErrInvalidJsonRequest struct {
	apperror.AppError
}

func NewErrInvalidJsonRequest() ErrInvalidJsonRequest {
	return ErrInvalidJsonRequest{apperror.New(
		errInvalidJsonRequestCode,
		errInvalidJsonRequestType,
		errInvalidJsonRequestMessage,
		http.StatusBadRequest,
		errInvalidJsonRequestDetails,
	)}
}

var (
	errEncodingResponseCode    = "A010003"
	errEncodingResponseType    = "InternalError"
	errEncodingResponseMessage = "unable to encode response"
	errEncodingResponseDetails = "unable to encode response"
)

type ErrEncodingResponse struct {
	apperror.AppError
}

func NewErrEncodingResponse() ErrEncodingResponse {
	return ErrEncodingResponse{apperror.New(
		errEncodingResponseCode,
		errEncodingResponseType,
		errEncodingResponseMessage,
		http.StatusInternalServerError,
		errEncodingResponseDetails,
	)}
}
