package dto

import (
	"tx-api/core/apperror"
	"tx-api/errors"

	"github.com/asaskevich/govalidator"
)

type CreateAccountRequest struct {
	DocumentNumber string `json:"documentNumber" valid:"required"`
}

func (req *CreateAccountRequest) Validate() apperror.Errors {
	var valErrors apperror.Errors
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		valErrors = req.mapValidatorErr2WpErr(err)
	}

	if len(valErrors) > 0 {
		return valErrors
	}
	return nil
}

func (req *CreateAccountRequest) mapValidatorErr2WpErr(err error) apperror.Errors {
	var wpErr apperror.Errors
	errs := err.(govalidator.Errors).Errors()
	for _, e := range errs {
		if gErr, ok := e.(govalidator.Errors); ok {
			wpErr = append(wpErr, req.mapValidatorErr2WpErr(gErr)...)
			continue
		}
		gErr := e.(govalidator.Error)
		wpErr = append(wpErr, errors.NewErrFieldMissing(gErr.Name).AppError)
	}
	return wpErr
}
