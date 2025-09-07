package errors

import (
	"fmt"
	"net/http"
	"tx-api/core/apperror"
)

var (
	errFieldMissingCode    = "A010001"
	errFieldMissingType    = "InvalidRequest"
	errFieldMissingMessage = "field missing/invalid in the request"
	errFieldMissingDetails = func(fieldName string) string {
		return fmt.Sprintf(
			"%s field missing or invalid in the request", fieldName)
	}
)

type ErrFieldMissing struct {
	apperror.AppError
}

func NewErrFieldMissingorInvalid(fieldName string) ErrFieldMissing {
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
	errEncodingResponseType    = "InternalServerError"
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

var (
	errDatabaseErrorCode    = "A010004"
	errDatabaseErrorType    = "InternalServerError"
	errDatabaseErrorMessage = "unable to encode response"
	errDatabaseErrorDetails = "unable to encode response"
)

type ErrDatabaseError struct {
	apperror.AppError
}

func NewErrDatabaseError() ErrDatabaseError {
	return ErrDatabaseError{apperror.New(
		errDatabaseErrorCode,
		errDatabaseErrorType,
		errDatabaseErrorMessage,
		http.StatusInternalServerError,
		errDatabaseErrorDetails,
	)}
}

var (
	errUniqueViolationCode    = "A010005"
	errUniqueViolationType    = "InternalServerError"
	errUniqueViolationMessage = "unique violation"
	errUniqueViolationDetails = func(fieldName, fieldValue string) string {
		return fmt.Sprintf(
			"%s as %s already present in system", fieldName, fieldValue)
	}
)

type ErrUniqueViolation struct {
	apperror.AppError
}

func NewErrUniqueViolation(fieldName, fieldValue string) ErrUniqueViolation {
	return ErrUniqueViolation{apperror.New(
		errUniqueViolationCode,
		errUniqueViolationType,
		errUniqueViolationMessage,
		http.StatusInternalServerError,
		errUniqueViolationDetails(fieldName, fieldValue),
	)}
}

var (
	errRecordNotFoundCode    = "A010006"
	errRecordNotFoundType    = "InternalServerError"
	errRecordNotFoundMessage = "record not found"
	errRecordNotFoundDetails = func(fieldName, fieldValue string) string {
		return fmt.Sprintf(
			"record not found for %s as %s in system", fieldName, fieldValue)
	}
)

type ErrRecordNotFound struct {
	apperror.AppError
}

func NewErrRecordNotFound(fieldName, fieldValue string) ErrRecordNotFound {
	return ErrRecordNotFound{apperror.New(
		errRecordNotFoundCode,
		errRecordNotFoundType,
		errRecordNotFoundMessage,
		http.StatusInternalServerError,
		errRecordNotFoundDetails(fieldName, fieldValue),
	)}
}

var (
	errForeignKeyViolationCode    = "A010007"
	errForeignKeyViolationType    = "InternalServerError"
	errForeignKeyViolationMessage = "foreign key violation"
	errForeignKeyViolationDetails = func(fieldName, fieldValue string) string {
		return fmt.Sprintf(
			"%s as %s does not exists in system", fieldName, fieldValue)
	}
)

type ErrForeignKeyViolation struct {
	apperror.AppError
}

func NewErrForeignKeyViolation(fieldName, fieldValue string) ErrForeignKeyViolation {
	return ErrForeignKeyViolation{apperror.New(
		errForeignKeyViolationCode,
		errForeignKeyViolationType,
		errForeignKeyViolationMessage,
		http.StatusInternalServerError,
		errForeignKeyViolationDetails(fieldName, fieldValue),
	)}
}
